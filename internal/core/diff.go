package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Diff compares two commits or the working directory with the latest commit
func Diff(repoPath, commitA, commitB string) error {
	gitDir := filepath.Join(repoPath, ".git")
	var err error

	if commitB == "" {
		// Compare working tree with commitA
		err = runDiffCommand(repoPath, commitA, "")
	} else {
		// Compare commitA and commitB
		err = runDiffCommand(gitDir, commitA, commitB)
	}

	if err != nil {
		return fmt.Errorf("diff failed: %w", err)
	}
	return nil
}

func runDiffCommand(repoPath, commitA, commitB string) error {
	cmd := exec.Command("diff", "-u", commitA, commitB)
	cmd.Dir = repoPath
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, errOut.String())
		return err
	}

	fmt.Println(out.String())
	return nil
}
