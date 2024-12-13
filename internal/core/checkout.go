package core

import (
	"fmt"
	"os"
	"path/filepath"
)

// Checkout switches to a branch or commit
func Checkout(repoPath, ref string) error {
	gitDir := filepath.Join(repoPath, ".git")

	// Check if the reference exists
	refPath := filepath.Join(gitDir, "refs", "heads", ref)
	if _, err := os.Stat(refPath); os.IsNotExist(err) {
		return fmt.Errorf("branch or commit %s does not exist", ref)
	}

	// Update HEAD
	headPath := filepath.Join(gitDir, "HEAD")
	if err := os.WriteFile(headPath, []byte(fmt.Sprintf("ref: refs/heads/%s\n", ref)), 0644); err != nil {
		return fmt.Errorf("failed to update HEAD: %w", err)
	}

	// Example: You can add logic to restore the working tree to the state of the ref
	fmt.Printf("Switched to branch '%s'\n", ref)
	return nil
}
