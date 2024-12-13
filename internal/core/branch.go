package core

import (
	"fmt"
	"os"
	"path/filepath"
)

// Branch lists or creates branches
func Branch(repoPath string, newBranch string) error {
	gitDir := filepath.Join(repoPath, ".git", "refs", "heads")

	if newBranch == "" {
		// List all branches
		branches, err := os.ReadDir(gitDir)
		if err != nil {
			return fmt.Errorf("failed to list branches: %w", err)
		}

		fmt.Println("Branches:")
		for _, branch := range branches {
			fmt.Println("*", branch.Name())
		}
		return nil
	}

	// Create a new branch
	branchPath := filepath.Join(gitDir, newBranch)
	headPath := filepath.Join(repoPath, ".git", "HEAD")
	headData, err := os.ReadFile(headPath)
	if err != nil {
		return fmt.Errorf("failed to read HEAD: %w", err)
	}

	if err := os.WriteFile(branchPath, headData, 0644); err != nil {
		return fmt.Errorf("failed to create branch: %w", err)
	}

	fmt.Printf("Branch '%s' created.\n", newBranch)
	return nil
}
