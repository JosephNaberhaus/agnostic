package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

var gitRootDir = ""

func GitRootDir() string {
	if len(gitRootDir) == 0 {
		output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
		if err != nil {
			panic(fmt.Errorf("failed to get the git root directory: %w", err))
		}

		gitRootDir = strings.TrimSpace(string(output))
	}

	return gitRootDir
}
