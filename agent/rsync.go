package agent

import (
	"os/exec"
)

var rsyncCommand = exec.Command

func Rsync(sourceDir, targetDir string, excludedFiles []string) error {
	arguments := append([]string{
		"--archive", "--delete",
		sourceDir + "/", targetDir,
	}, makeExclusionList(excludedFiles)...)

	_, err := rsyncCommand("rsync", arguments...).Output()

	return err
}

func makeExclusionList(excludedFiles []string) []string {
	var exclusions []string
	for _, excludedFile := range excludedFiles {
		exclusions = append(exclusions, "--exclude", excludedFile)
	}
	return exclusions
}
