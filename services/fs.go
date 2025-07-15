package services

import (
	"os"
	"path/filepath"
	"strings"
)

// ReadFile reads content from a file
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WriteFile writes content to a file
func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

// DeleteFile removes a file
func DeleteFile(path string) error {
	return os.Remove(path)
}

// ListFiles returns all files with given extension in directory
func ListFiles(dir string, ext string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ext) {
			fullPath := filepath.Join(dir, entry.Name())
			files = append(files, fullPath)
		}
	}

	return files, nil
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
