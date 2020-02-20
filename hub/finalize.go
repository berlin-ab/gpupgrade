package hub

import (
	"fmt"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/hashicorp/go-multierror"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/step"
)

func (s *Server) Finalize(_ *idl.FinalizeRequest, stream idl.CliToHub_FinalizeServer) (err error) {
	agentConnections, err := s.AgentConns()
	if err != nil {
		return err
	}

	st, err := BeginStep(s.StateDir, "finalize", stream)
	if err != nil {
		return err
	}

	defer func() {
		if ferr := st.Finish(); ferr != nil {
			err = multierror.Append(err, ferr).ErrorOrNil()
		}

		if err != nil {
			gplog.Error(fmt.Sprintf("finalize: %s", err))
		}
	}()

	st.Run(idl.Substep_FINALIZE_SHUTDOWN_TARGET_CLUSTER, func(streams step.OutStreams) error {
		return StopCluster(streams, s.Target, false)
	})

	st.Run(idl.Substep_FINALIZE_SWAP_DATA_DIRECTORIES, func(streams step.OutStreams) error {
		agentBroker := &AgentBrokerGRPC{
			agentConnections: agentConnections,
		}
		hub := MakeHub(s.Config)

		return SwapDataDirectories(hub, agentBroker)
	})

	st.Run(idl.Substep_FINALIZE_START_TARGET_MASTER, func(streams step.OutStreams) error {
		return StartTargetMasterForFinalize(s.Config, streams)
	})

	// Once UpdateCatalog && UpdateMasterConf is executed, the port on which the target
	// cluster starts is changed in the catalog and postgresql.conf, however the server config.json target port is
	// still the old port on which the target cluster was initialized.
	// TODO: if any steps needs to connect to the new cluster (that should use new port), we should either
	// write it to the config.json or add some way to identify the state.
	st.Run(idl.Substep_FINALIZE_UPDATE_CATALOG, func(streams step.OutStreams) error {
		return UpdateCatalog(s.Source, s.Target)
	})

	st.Run(idl.Substep_FINALIZE_SHUTDOWN_TARGET_MASTER, func(streams step.OutStreams) error {
		return StopMasterOnly(streams, s.Target, false)
	})

	st.Run(idl.Substep_FINALIZE_UPDATE_POSTGRESQL_CONF, func(streams step.OutStreams) error {
		return UpdateMasterConf(s.Source, s.Target)
	})

	st.Run(idl.Substep_FINALIZE_START_TARGET_CLUSTER, func(streams step.OutStreams) error {
		return StartCluster(streams, s.Target, false)
	})

	return st.Err()
}

func StartTargetMasterForFinalize(config *Config, streams step.OutStreams) error {
	// We need to pass a modified Target cluster to StartMasterOnly because
	// the data directory has been promoted to its new location
	var target = *config.Target
	targetMaster := target.Primaries[-1]
	targetMaster.DataDir = targetMaster.PromotionDataDirectory(config.Source.Primaries[-1])
	target.Primaries[-1] = targetMaster

	return StartMasterOnly(streams, &target, false)
}
