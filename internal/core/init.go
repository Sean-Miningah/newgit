package core

import (
	"fmt"
	"os"
	"path/filepath"
)

// Initialize a new repo
func InitRepo(path string) error {
	gitDir := filepath.Join(path, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		return fmt.Errorf("repository already exists in %s", path)
	}

	// Create Necessary directories for version control
	dirs := []string{
		"branches",
		"objects",
		"refs/heads",
		"refs/tags",
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(path, dir), 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", dir, err)
		}
	}

	// Create default files
	files := map[string]string{
		"description": "Unnamed repository; edit this file to name the repository.\n",
		"HEAD":        "ref: refs/heads/master\n",
		"config": `[core]
			repositoryformatversion = 0
			filemode = false
			bare = false
			`,
	}
	for name, content := range files {
		err := os.WriteFile(filepath.Join(gitDir, name), []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("failed to write %s: %w", name, err)
		}
	}

	fmt.Printf("Initialized empty Git repository in %s\n", gitDir)
	return nil
}
