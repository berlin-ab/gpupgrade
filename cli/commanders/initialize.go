package commanders

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/greenplum-db/gpupgrade/hub/state"
	"github.com/greenplum-db/gpupgrade/idl"

	"github.com/greenplum-db/gp-common-go-libs/gplog"

	"github.com/greenplum-db/gpupgrade/utils"
)

// introduce this variable to allow exec.Command to be mocked out in tests
var execCommandHubStart = exec.Command
var execCommandHubCount = exec.Command

// we create the state directory in the cli to ensure that at most one gpupgrade is occurring
// at the same time.
func CreateStateDir() (err error) {
	s := Substep(idl.Substep_CREATING_DIRECTORIES)
	defer s.Finish(&err)

	stateDir := utils.GetStateDir()
	err = os.Mkdir(stateDir, 0700)
	if os.IsExist(err) {
		gplog.Debug("Config directory %s already present...skipping", stateDir)
		return nil
	}
	if err != nil {
		gplog.Debug("Config directory %s could not be created.", stateDir)
		return err
	}

	return nil
}

// TODO: move the creation of this file to the hub
func createConfigFile() error {
	// if empty json configuration file exists, skip recreating it
	filepath := state.GetConfigFilepath(utils.GetStateDir())
	_, err := os.Stat(filepath)

	// if the file exists, there will be no error or if there is an error it might
	// also indicate that the file exists, in either case don't overwrite the file
	if err == nil || os.IsExist(err) {
		gplog.Debug("Initial cluster configuration file %s already present...skipping", filepath)
		return nil
	}

	// if the err is anything other than file does not exist, error out
	if !os.IsNotExist(err) {
		gplog.Debug("Check to find presence of initial cluster configuration file %s failed", filepath)
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "{}") // the hub will fill this in during initialization

	return nil
}

func CreateInitialClusterConfigs() (err error) {
	s := Substep(idl.Substep_GENERATING_CONFIG)
	defer s.Finish(&err)

	err = createConfigFile()

	if err != nil {
		return err
	}

	return nil
}

func StartHub() (err error) {
	s := Substep(idl.Substep_START_HUB)
	defer s.Finish(&err)

	running, err := IsHubRunning()
	if err != nil {
		gplog.Error("failed to determine if hub already running")
		return err
	}
	if running {
		gplog.Debug("gpupgrade hub already running...skipping.")
		return nil
	}

	cmd := execCommandHubStart("gpupgrade", "hub", "--daemonize")
	stdout, cmdErr := cmd.Output()
	if cmdErr != nil {
		err := fmt.Errorf("failed to start hub (%s)", cmdErr)
		if exitErr, ok := cmdErr.(*exec.ExitError); ok {
			// Annotate with the Stderr capture, if we have it.
			err = fmt.Errorf("%s: %s", err, exitErr.Stderr)
		}
		return err
	}
	gplog.Debug("gpupgrade hub started successfully: %s", stdout)
	return nil
}

func IsHubRunning() (bool, error) {
	script := `ps -ef | grep -wGc "[g]pupgrade hub"` // use square brackets to avoid finding yourself in matches
	_, err := execCommandHubCount("bash", "-c", script).Output()

	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ProcessState.ExitCode() == 1 { // hub not found
			return false, nil
		}
	}
	if err != nil { // grep failed
		return false, err
	}

	return true, nil
}
