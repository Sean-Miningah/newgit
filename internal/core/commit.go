package core

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Commit(repoPath, message string) error {
	gitDir := filepath.Join(repoPath, ".git")
	objectsDir := filepath.Join(gitDir, "objects")

	// Generate commit content
	content := fmt.Sprintf("commit\n%s\n%s\n%s\n", time.Now().String(), message, "tree_hash_placeholder")

	// Compute hash and write object
	hash := ComputeHash(content)
	commitPath := filepath.Join(objectsDir, hash[:2], hash[2:])
	if err := WriteObject(commitPath, content); err != nil {
		return err
	}

	// Update HEAD
	headPath := filepath.Join(gitDir, "HEAD")
	return os.WriteFile(headPath, []byte(fmt.Sprintf("ref: refs/heads/master\n%s\n", hash)), 0644)
}

// ComputeHash computes the SHA-1 hash of a string
func ComputeHash(content string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(content)))
}

// WriteObject writes a Git object to the filesystem
func WriteObject(path, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}
	return os.WriteFile(path, []byte(content), 0644)
}
