package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Log(repoPath string) error {
	headPath := filepath.Join(repoPath, ".git", "HEAD")
	headData, err := os.ReadFile(headPath)

	if err != nil {
		return fmt.Errorf("failed to read HEAD: %w", err)
	}

	ref := strings.TrimSpace(string(headData))
	if strings.HasPrefix(ref, "ref:") {
		refPath := filepath.Join(repoPath, ".git", ref[4:])
		refData, err := os.ReadFile(refPath)
		if err != nil {
			return fmt.Errorf("failed to read ref: %w", err)
		}

		ref = strings.TrimSpace(string(refData))
	}

	for ref != "" {
		commitPath := filepath.Join(repoPath, ".git", "objects", ref[:2], ref[2:])
		commitData, err := os.ReadFile(commitPath)
		if err != nil {
			return fmt.Errorf("failed to read commit object: %w", err)
		}

		// Example parsing commit data
		fmt.Println("Commit:", ref)
		fmt.Println(string(commitData))

		ref = extractParentCommit(string(commitData))
	}

	return nil
}

func extractParentCommit(commitData string) string {
	lines := strings.Split(commitData, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "parent ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "parent "))
		}
	}
	return ""
}
