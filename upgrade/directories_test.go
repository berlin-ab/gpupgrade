// Copyright (c) 2017-2020 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package upgrade_test

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"testing"

	"github.com/greenplum-db/gp-common-go-libs/testhelper"
	"golang.org/x/xerrors"

	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/go-multierror"

	"github.com/greenplum-db/gpupgrade/testutils"
	"github.com/greenplum-db/gpupgrade/upgrade"
	"github.com/greenplum-db/gpupgrade/utils"
)

func TestTempDataDir(t *testing.T) {
	var id upgrade.ID

	cases := []struct {
		datadir        string
		segPrefix      string
		expectedFormat string // %s will be replaced with id.String()
	}{
		{"/data/seg-1", "seg", "/data/seg.%s.-1"},
		{"/data/master/gpseg-1", "gpseg", "/data/master/gpseg.%s.-1"},
		{"/data/seg1", "seg", "/data/seg.%s.1"},
		{"/data/seg1/", "seg", "/data/seg.%s.1"},
		{"/data/standby", "seg", "/data/standby.%s"},
	}

	for _, c := range cases {
		actual := upgrade.TempDataDir(c.datadir, c.segPrefix, id)
		expected := fmt.Sprintf(c.expectedFormat, id)

		if actual != expected {
			t.Errorf("TempDataDir(%q, %q, id) = %q, want %q",
				c.datadir, c.segPrefix, actual, expected)
		}
	}
}

func ExampleTempDataDir() {
	var id upgrade.ID

	master := upgrade.TempDataDir("/data/master/seg-1", "seg", id)
	standby := upgrade.TempDataDir("/data/standby", "seg", id)
	segment := upgrade.TempDataDir("/data/primary/seg3", "seg", id)

	fmt.Println(master)
	fmt.Println(standby)
	fmt.Println(segment)
	// Output:
	// /data/master/seg.AAAAAAAAAAA.-1
	// /data/standby.AAAAAAAAAAA
	// /data/primary/seg.AAAAAAAAAAA.3
}

func TestRenameDataDirectory(t *testing.T) {
	testhelper.SetupTestLogger()

	t.Run("successfully renames source to archive, and target to source", func(t *testing.T) {
		source, archive, target, cleanup := mustCreateDirs(t)
		defer cleanup(t)

		err := upgrade.RenameDataDirectory(source, archive, target, true)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		verifyRename(t, source, archive, target)
	})

	t.Run("returns early if already renamed", func(t *testing.T) {
		source := testutils.GetTempDir(t, "source")
		defer os.RemoveAll(source)

		archive := testutils.GetTempDir(t, "archive")
		defer os.RemoveAll(archive)

		target := ""

		called := false
		utils.System.Rename = func(old, new string) error {
			called = true
			return nil
		}
		defer func() {
			utils.System.Rename = os.Rename
		}()

		verifyRename(t, source, archive, target)

		err := upgrade.RenameDataDirectory(source, archive, target, true)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		if called {
			t.Errorf("expected rename to not be called")
		}
	})

	t.Run("bubbles up errors", func(t *testing.T) {
		source, archive, target, cleanup := mustCreateDirs(t)
		defer cleanup(t)

		expected := errors.New("permission denied")
		utils.System.Rename = func(old, new string) error {
			return expected
		}
		defer func() {
			utils.System.Rename = os.Rename
		}()

		err := upgrade.RenameDataDirectory(source, archive, target, true)
		if !xerrors.Is(err, expected) {
			t.Errorf("got %#v want %#v", err, expected)
		}
	})

	t.Run("it returns other LinkErrors when renaming the source fails for errors other than ENOENT", func(t *testing.T) {
		source, archive, target, cleanup := mustCreateDirs(t)
		defer cleanup(t)

		expected := &os.LinkError{Err: syscall.EEXIST}
		utils.System.Rename = func(old, new string) error {
			if old == source {
				return expected
			}
			return os.Rename(old, new)
		}
		defer func() {
			utils.System.Rename = os.Rename
		}()

		err := upgrade.RenameDataDirectory(source, archive, target, true)
		if !xerrors.Is(err, expected) {
			t.Errorf("got %#v want %#v", err, expected)
		}
	})

	t.Run("only renames source to archive when renameTarget is false", func(t *testing.T) {
		source := testutils.GetTempDir(t, "source-")
		defer os.RemoveAll(source)

		archive := source + upgrade.OldSuffix

		calls := 0
		utils.System.Rename = func(old, new string) error {
			calls++

			if old != source {
				t.Errorf("got %q want %q", old, source)
			}

			if new != archive {
				t.Errorf("got %q want %q", new, archive)
			}

			return os.Rename(old, new)
		}
		defer func() {
			utils.System.Rename = os.Rename
		}()

		err := upgrade.RenameDataDirectory(source, archive, "", false)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		if calls != 1 {
			t.Errorf("expected rename to be called once")
		}

		if upgrade.PathExists(source) {
			t.Errorf("expected source %q to not exist", source)
		}

		if !upgrade.PathExists(archive) {
			t.Errorf("expected archive %q to exist", archive)
		}
	})

	t.Run("when renaming succeeds then a re-run succeeds", func(t *testing.T) {
		source, archive, target, cleanup := mustCreateDirs(t)
		defer cleanup(t)

		err := upgrade.RenameDataDirectory(source, archive, target, true)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		verifyRename(t, source, archive, target)

		err = upgrade.RenameDataDirectory(source, archive, target, true)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		verifyRename(t, source, archive, target)
	})

	t.Run("when renaming the source fails then a re-run succeeds", func(t *testing.T) {
		source, archive, target, cleanup := mustCreateDirs(t)
		defer cleanup(t)

		expected := errors.New("permission denied")
		utils.System.Rename = func(old, new string) error {
			if old == source {
				return expected
			}
			return os.Rename(old, new)
		}

		err := upgrade.RenameDataDirectory(source, archive, target, true)
		if !xerrors.Is(err, expected) {
			t.Errorf("got %#v want %#v", err, expected)
		}

		if !upgrade.PathExists(source) {
			t.Errorf("expected source %q to exist", source)
		}

		if upgrade.PathExists(archive) {
			t.Errorf("expected archive %q to not exist", archive)
		}

		if !upgrade.PathExists(target) {
			t.Errorf("expected target %q to exist", target)
		}

		utils.System.Rename = os.Rename

		err = upgrade.RenameDataDirectory(source, archive, target, true)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		verifyRename(t, source, archive, target)
	})

	t.Run("when renaming the target fails then a re-run succeeds", func(t *testing.T) {
		source, archive, target, cleanup := mustCreateDirs(t)
		defer cleanup(t)

		expected := errors.New("permission denied")
		utils.System.Rename = func(old, new string) error {
			if old == target {
				return expected
			}
			return os.Rename(old, new)
		}

		err := upgrade.RenameDataDirectory(source, archive, target, true)
		if !xerrors.Is(err, expected) {
			t.Errorf("got %#v want %#v", err, expected)
		}

		if upgrade.PathExists(source) {
			t.Errorf("expected source %q to not exist", source)
		}

		if !upgrade.PathExists(archive) {
			t.Errorf("expected archive %q to exist", archive)
		}

		if !upgrade.PathExists(target) {
			t.Errorf("expected target %q to exist", target)
		}

		utils.System.Rename = os.Rename

		err = upgrade.RenameDataDirectory(source, archive, target, true)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		verifyRename(t, source, archive, target)
	})
}

func mustCreateDirs(t *testing.T) (string, string, string, func(*testing.T)) {
	t.Helper()

	source := testutils.GetTempDir(t, "source")
	archive := source + upgrade.OldSuffix
	target := testutils.GetTempDir(t, "target")

	return source, archive, target, func(t *testing.T) {
		if err := os.RemoveAll(source); err != nil {
			t.Errorf("removing source directory: %v", err)
		}
		if err := os.RemoveAll(archive); err != nil {
			t.Errorf("removing archive directory: %v", err)
		}
		if err := os.RemoveAll(target); err != nil {
			t.Errorf("removing target directory: %v", err)
		}
	}
}

func verifyRename(t *testing.T, source, archive, target string) {
	t.Helper()

	if !upgrade.PathExists(source) {
		t.Errorf("expected source %q to exist", source)
	}

	if !upgrade.PathExists(archive) {
		t.Errorf("expected archive %q to exist", archive)
	}

	if upgrade.PathExists(target) {
		t.Errorf("expected target %q to not exist", target)
	}
}

func setup(t *testing.T) (teardown func(), directories []string, requiredPaths []string) {
	requiredPaths = []string{"pg_file1", "pg_file2"}
	var dataDirectories = []string{"/data/dbfast_mirror1/seg1", "/data/dbfast_mirror2/seg2"}
	rootDir, directories := setupDirs(t, dataDirectories, requiredPaths)
	teardown = func() {
		err := os.RemoveAll(rootDir)
		if err != nil {
			t.Fatalf("error %#v when deleting directory %#v", err, rootDir)
		}
	}

	return teardown, directories, requiredPaths
}

func TestDeleteDirectories(t *testing.T) {
	testhelper.SetupTestLogger()

	t.Run("successfully deletes the directories if all required paths exist in that directory", func(t *testing.T) {
		teardown, directories, requiredPaths := setup(t)
		defer teardown()

		err := upgrade.DeleteDirectories(directories, requiredPaths)

		if err != nil {
			t.Errorf("unexpected error got %+v", err)
		}

		for _, dataDir := range directories {
			if _, err := os.Stat(dataDir); err == nil {
				t.Errorf("dataDir %s exists", dataDir)
			}
		}
	})

	t.Run("fails when the required paths are not in the directories", func(t *testing.T) {
		teardown, directories, _ := setup(t)
		defer teardown()

		err := upgrade.DeleteDirectories(directories, []string{"a", "b"})

		var multiErr *multierror.Error
		if !xerrors.As(err, &multiErr) {
			t.Fatalf("got error %#v, want type %T", err, multiErr)
		}

		if len(multiErr.Errors) != 4 {
			t.Errorf("received %d errors, want %d", len(multiErr.Errors), 4)
		}

		for _, err := range multiErr.Errors {
			if !xerrors.Is(err, os.ErrNotExist) {
				t.Errorf("got error %#v, want %#v", err, os.ErrNotExist)
			}
		}
	})

	t.Run("fails to remove one segment data directory", func(t *testing.T) {
		teardown, directories, requiredPaths := setup(t)
		defer teardown()

		fileToRemove := filepath.Join(directories[0], requiredPaths[0])
		if err := os.Remove(fileToRemove); err != nil {
			t.Errorf("unexpected error %+v", err)
		}

		err2 := upgrade.DeleteDirectories(directories, requiredPaths)

		var multiErr *multierror.Error
		if !xerrors.As(err2, &multiErr) {
			t.Fatalf("got error %#v, want type %T", err2, multiErr)
		}

		if len(multiErr.Errors) != 1 {
			t.Errorf("got %d errors, want %d", len(multiErr.Errors), 1)
		}

		var actualErr *os.PathError

		for _, err := range multiErr.Errors {
			if !xerrors.As(err, &actualErr) {
				t.Errorf("got error %#v, want %#v", err, "PathError")
			}
		}

		if _, err := os.Stat(directories[0]); err != nil {
			t.Errorf("dataDir should exist, stat error %+v", err)
		}

		if _, err := os.Stat(directories[1]); err == nil {
			t.Errorf("dataDir %s exists", directories[1])
		}
	})
}

func setupDirs(t *testing.T, subdirectories []string, requiredPaths []string) (tmpDir string, createdDirectories []string) {
	var err error
	tmpDir, err = ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("error creating temporary directory: %v", err)
	}

	for _, dir := range subdirectories {
		createdDirectories = append(createdDirectories, createDataDir(t, dir, tmpDir, requiredPaths))
	}

	return tmpDir, createdDirectories
}

func createDataDir(t *testing.T, name, tmpDir string, requiredPaths []string) (dirPath string) {
	dirPath = filepath.Join(tmpDir, name)

	err := os.MkdirAll(dirPath, 0700)
	if err != nil {
		t.Errorf("error creating path: %v", err)
	}

	for _, fileName := range requiredPaths {
		filePath := filepath.Join(dirPath, fileName)
		err = ioutil.WriteFile(filePath, []byte{}, 0600)
		if err != nil {
			t.Errorf("error writing empty file: %v", err)
		}
	}

	return dirPath
}