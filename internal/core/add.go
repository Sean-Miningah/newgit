package core

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func AddFile(repoPath, filePath string) error {
	objectDir := filepath.Join(repoPath, ".git", "objects")
	if err := os.MkdirAll(objectDir, 0755); err != nil {
		return fmt.Errorf("failed to create objects directory: %w", err)
	}

	// Read file contents
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Compute SHA-1 hash
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return fmt.Errorf("failed to hash file contents: %w", err)
	}

	sha := fmt.Sprintf("%x", hash.Sum(nil))

	// Write object file
	objectPath := filepath.Join(objectDir, sha[:2], sha[2:])
	if err := os.MkdirAll(filepath.Dir(objectPath), 0755); err != nil {
		return fmt.Errorf("failed to create object directory: %w", err)
	}
	file.Seek(0, io.SeekStart)
	out, err := os.Create(objectPath)
	if err != nil {
		return fmt.Errorf("failed to create object file: %w", err)
	}
	defer out.Close()
	if _, err := io.Copy(out, file); err != nil {
		return fmt.Errorf("failed to write object file: %w", err)
	}

	fmt.Printf("Added file %s as %s\n", filePath, sha)
	return nil
}
