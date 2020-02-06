package agent_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"golang.org/x/xerrors"

	"github.com/greenplum-db/gpupgrade/agent"
	"github.com/greenplum-db/gpupgrade/testutils/exectest"
)

func getTempDir(t *testing.T) string {
	sourceDir, err := ioutil.TempDir("", "rsync-source")
	if err != nil {
		t.Fatalf("creating temporary directory: %+v", err)
	}

	return sourceDir
}

func writeToFile(filepath string, contents []byte, t *testing.T) {
	err := ioutil.WriteFile(filepath, contents, 0644)

	if err != nil {
		t.Fatalf("error writing file '%v'", filepath)
	}
}

func TestRsync(t *testing.T) {
	if _, err := exec.LookPath("rsync"); err != nil {
		t.Skipf("tests require rsync (%v)", err)
	}

	// These are "live" integration tests. Plug exec.Command back into the
	// system.
	agent.SetRsyncCommand(exec.Command)
	defer func() { agent.SetRsyncCommand(nil) }()

	t.Run("it copies data from a source directory to a target directory", func(t *testing.T) {
		sourceDir := getTempDir(t)
		defer os.RemoveAll(sourceDir)

		targetDir := getTempDir(t)
		defer os.RemoveAll(targetDir)

		writeToFile(sourceDir+"/hi", []byte("hi"), t)

		if err := agent.Rsync(sourceDir, targetDir, []string{}); err != nil {
			t.Errorf("Rsync() returned error %+v", err)
		}

		targetContents, _ := ioutil.ReadFile(filepath.Join(targetDir, "/hi"))

		if bytes.Compare(targetContents, []byte("hi")) != 0 {
			t.Errorf("target directory file 'hi' contained %v, wanted %v",
				targetContents,
				"hi")
		}
	})

	t.Run("it removes files that existed in the target directory before the sync", func(t *testing.T) {
		sourceDir := getTempDir(t)
		defer os.RemoveAll(sourceDir)

		targetDir := getTempDir(t)
		defer os.RemoveAll(targetDir)

		writeToFile(targetDir+"/file-that-should-get-removed", []byte("goodbye"), t)

		if err := agent.Rsync(sourceDir, targetDir, []string{}); err != nil {
			t.Errorf("Rsync() returned error %+v", err)
		}

		_, statError := os.Stat(filepath.Join(targetDir, "file-that-should-get-removed"))

		if os.IsExist(statError) {
			t.Errorf("target directory file 'file-that-should-get-removed' should not exist, but it does")
		}
	})

	t.Run("it does not copy files from the source directory when in the exclusion list", func(t *testing.T) {
		sourceDir := getTempDir(t)
		defer os.RemoveAll(sourceDir)

		targetDir := getTempDir(t)
		defer os.RemoveAll(targetDir)

		writeToFile(filepath.Join(sourceDir, "file-that-should-get-excluded"), []byte("goodbye"), t)

		err := agent.Rsync(sourceDir, targetDir, []string{"file-that-should-get-excluded"})
		if err != nil {
			t.Errorf("Rsync() returned error %+v", err)
		}

		_, statError := os.Stat(filepath.Join(targetDir, "file-that-should-get-excluded"))

		if os.IsExist(statError) {
			t.Errorf("target directory file 'file-that-should-get-excluded' should not exist, but it does")
		}
	})

	t.Run("it preserves files in the target directory when in the exclusion list", func(t *testing.T) {
		sourceDir := getTempDir(t)
		defer os.RemoveAll(sourceDir)

		targetDir := getTempDir(t)
		defer os.RemoveAll(targetDir)

		writeToFile(filepath.Join(sourceDir, "file-that-should-get-copied"), []byte("new file"), t)
		writeToFile(filepath.Join(targetDir, "file-that-should-get-ignored"), []byte("i'm still here"), t)
		writeToFile(filepath.Join(targetDir, "another-file-that-should-get-ignored"), []byte("i'm still here"), t)

		err := agent.Rsync(sourceDir, targetDir, []string{"file-that-should-get-ignored", "another-file-that-should-get-ignored"})
		if err != nil {
			t.Errorf("Rsync() returned error %+v", err)
		}

		_, statError := os.Stat(filepath.Join(targetDir, "file-that-should-get-ignored"))

		if os.IsNotExist(statError) {
			t.Error("target directory file 'file-that-should-get-ignored' should still exist, but it does not")
		}

		_, statError = os.Stat(filepath.Join(targetDir, "another-file-that-should-get-ignored"))

		if os.IsNotExist(statError) {
			t.Error("target directory file 'another-file-that-should-get-ignored' should still exist, but it does not")
		}

		_, statError = os.Stat(filepath.Join(targetDir, "file-that-should-get-copied"))

		if os.IsNotExist(statError) {
			t.Error("target directory file 'file-that-should-get-copied' should exist, but does not")
		}
	})

	t.Run("it bubbles up errors", func(t *testing.T) {
		agent.SetRsyncCommand(exectest.NewCommand(agent.FailedRsync))

		sourceDir := getTempDir(t)
		defer os.RemoveAll(sourceDir)

		targetDir := "/tmp/some/invalid/target/dir"
		defer os.RemoveAll(targetDir)

		writeToFile(filepath.Join(sourceDir, "some-file"), []byte("hi"), t)

		copyError := agent.Rsync(sourceDir, targetDir, []string{""})

		if copyError == nil {
			t.Errorf("got no copy errors, wanted a copy error because target directory did not exist")
		}

		var exitError *exec.ExitError

		if !xerrors.As(copyError, &exitError) {
			t.Errorf("got an error type %T, want %T",
				copyError,
				exitError)
		}

		if string(exitError.Stderr) != "rsync failed cause I said so" {
			t.Errorf("expected error 'rsync failed cause I said so', got %q", string(exitError.Stderr))
		}

	})
}
