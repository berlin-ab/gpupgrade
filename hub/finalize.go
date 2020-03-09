package hub

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"golang.org/x/xerrors"

	"github.com/greenplum-db/gp-common-go-libs/gplog"

	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/step"
)

func (s *Server) Finalize(_ *idl.FinalizeRequest, stream idl.CliToHub_FinalizeServer) (err error) {
	st, err := step.Begin(s.StateDir, "finalize", stream)
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

	if s.Source.HasStandby() {
		st.Run(idl.Substep_FINALIZE_UPGRADE_STANDBY, func() error {
			greenplumRunner := &greenplumRunner{
				masterPort:          s.Target.MasterPort(),
				masterDataDirectory: s.Target.MasterDataDir(),
				binDir:              s.Target.BinDir,
				streams:             st.Streams(),
			}

			// TODO: Persist the standby to config.json and update the
			//  source & target clusters.
			return UpgradeStandby(greenplumRunner, StandbyConfig{
				Port:          s.TargetInitializeConfig.Standby.Port,
				Hostname:      s.Source.StandbyHostname(),
				DataDirectory: s.Source.StandbyDataDirectory() + "_upgrade",
			})
		})
	}

	st.Run(idl.Substep_FINALIZE_SHUTDOWN_TARGET_CLUSTER, func() error {
		err := StopCluster(st.Streams(), s.Target)

		if err != nil {
			return xerrors.Errorf("failed to stop target cluster: %w", err)
		}

		return nil
	})

	st.Run(idl.Substep_FINALIZE_UPDATE_TARGET_CATALOG_AND_CLUSTER_CONFIG, func() error {
		return s.UpdateCatalogAndClusterConfig(st.Streams())
	})

	st.Run(idl.Substep_FINALIZE_RENAME_DATA_DIRECTORIES, func() error {
		return s.RenameDataDirectories()
	})

	st.Run(idl.Substep_FINALIZE_UPDATE_TARGET_CONF_FILES, func() error {
		return s.UpdateConfFiles()
	})

	st.Run(idl.Substep_FINALIZE_START_TARGET_CLUSTER, func() error {
		err := StartCluster(st.Streams(), s.Target)

		if err != nil {
			return xerrors.Errorf("failed to start target cluster: %w", err)
		}

		return nil
	})

	return st.Err()
}
