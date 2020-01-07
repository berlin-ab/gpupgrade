package hub

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/greenplum-db/gp-common-go-libs/cluster"
	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/greenplum-db/gpupgrade/db"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/step"
	"github.com/greenplum-db/gpupgrade/utils"
)

func (h *Hub) Initialize(in *idl.InitializeRequest, stream idl.CliToHub_InitializeServer) (err error) {
	s, err := BeginStep(h.conf.StateDir, "initialize", stream)
	if err != nil {
		return err
	}

	defer func() {
		if ferr := s.Finish(); ferr != nil {
			err = multierror.Append(err, ferr).ErrorOrNil()
		}

		if err != nil {
			gplog.Error(fmt.Sprintf("initialize: %s", err))
		}
	}()

	s.Run(idl.UpgradeSteps_CONFIG, func(stream step.OutStreams) error {
		return h.fillClusterConfigsSubStep(stream, in.OldBinDir, in.NewBinDir, int(in.OldPort))
	})

	s.Run(idl.UpgradeSteps_START_AGENTS, func(stream step.OutStreams) error {
		return h.startAgentsSubStep(stream)
	})

	return s.Err()
}

func (h *Hub) InitializeCreateCluster(in *idl.InitializeCreateClusterRequest, stream idl.CliToHub_InitializeCreateClusterServer) (err error) {
	s, err := BeginStep(h.conf.StateDir, "initialize", stream)
	if err != nil {
		return err
	}

	defer func() {
		if ferr := s.Finish(); ferr != nil {
			err = multierror.Append(err, ferr).ErrorOrNil()
		}

		if err != nil {
			gplog.Error(fmt.Sprintf("initialize: %s", err))
		}
	}()

	var targetMasterPort int
	s.Run(idl.UpgradeSteps_CREATE_TARGET_CONFIG, func(_ step.OutStreams) error {
		var err error
		targetMasterPort, err = h.GenerateInitsystemConfig(in.Ports)
		return err
	})

	s.Run(idl.UpgradeSteps_SHUTDOWN_SOURCE_CLUSTER, func(stream step.OutStreams) error {
		return StopCluster(stream, h.source)
	})

	s.Run(idl.UpgradeSteps_INIT_TARGET_CLUSTER, func(stream step.OutStreams) error {
		return h.CreateTargetCluster(stream, targetMasterPort)
	})

	s.Run(idl.UpgradeSteps_SHUTDOWN_TARGET_CLUSTER, func(stream step.OutStreams) error {
		return h.ShutdownCluster(stream, false)
	})

	s.Run(idl.UpgradeSteps_CHECK_UPGRADE, func(stream step.OutStreams) error {
		return h.CheckUpgrade(stream)
	})

	return s.Err()
}

// create old/new clusters, write to disk and re-read from disk to make sure it is "durable"
func (h *Hub) fillClusterConfigsSubStep(_ OutStreams, oldBinDir, newBinDir string, oldPort int) error {
	conn := db.NewDBConn("localhost", oldPort, "template1")
	defer conn.Close()

	var err error
	h.Source, err = utils.ClusterFromDB(conn, oldBinDir)
	if err != nil {
		return errors.Wrap(err, "could not retrieve source configuration")
	}

	h.Target = &utils.Cluster{Cluster: new(cluster.Cluster), BinDir: newBinDir}

	if err := h.SaveConfig(); err != nil {
		return err
	}

	// link in source/target to hub
	// TODO: remove once we deduplicate
	h.source = h.Source
	h.target = h.Target

	return nil
}

func getAgentPath() (string, error) {
	hubPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(filepath.Dir(hubPath), "gpupgrade"), nil
}

// TODO: use the implementation in RestartAgents() for this function and combine them
func (h *Hub) startAgentsSubStep(stream OutStreams) error {
	source := h.source
	stateDir := h.conf.StateDir

	// XXX If there are failures, does it matter what agents have successfully
	// started, or do we just want to stop all of them and kick back to the
	// user?
	logStr := "start agents on master and hosts"

	agentPath, err := getAgentPath()
	if err != nil {
		return errors.Errorf("failed to get the hub executable path %v", err)
	}

	// XXX State directory handling on agents needs to be improved. See issue
	// #127: all agents will silently recreate that directory if it doesn't
	// already exist. Plus, ExecuteOnAllHosts() doesn't let us control whether
	// we execute locally or via SSH for the master, so we don't know whether
	// GPUPGRADE_HOME is going to be inherited.
	runAgentCmd := func(contentID int) string {
		return agentPath + " agent --daemonize --state-directory " + stateDir
	}

	errStr := "Failed to start all gpupgrade agents"

	remoteOutput, err := source.ExecuteOnAllHosts(logStr, runAgentCmd)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	errMessage := func(contentID int) string {
		return fmt.Sprintf("Could not start gpupgrade agent on segment with contentID %d", contentID)
	}
	source.CheckClusterError(remoteOutput, errStr, errMessage, true)

	// Agents print their port and PID to stdout; log them for posterity.
	for content, output := range remoteOutput.Stdouts {
		if remoteOutput.Errors[content] == nil {
			gplog.Info("[%s] %s", source.Segments[content].Hostname, output)
		}
	}

	if remoteOutput.NumErrors > 0 {
		// CheckClusterError() will have already logged each error.
		return errors.New("could not start agents on segment hosts; see log for details")
	}

	return nil
}