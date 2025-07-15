package note

import (
	"os"
	"path/filepath"
	"strings"
)

type Note struct {
	Name    string // file name without .md extension
	Path    string // full file path
	Content string // markdown content
}

// Load reads the note content from file
func (n *Note) Load() error {
	data, err := os.ReadFile(n.Path)
	if err != nil {
		return err
	}

	n.Content = string(data)
	return nil
}

// Save writes the note content to file
func (n *Note) Save() error {
	return os.WriteFile(n.Path, []byte(n.Content), 0644)
}

// Delete removes the note file
func (n *Note) Delete() error {
	return os.Remove(n.Path)
}

// ListNotes returns all notes from a directory
func ListNotes(dir string) ([]*Note, error) {
	files, err := ListFiles(dir, ".md")
	if err != nil {
		return nil, err
	}

	var notes []*Note
	for _, filePath := range files {
		fileName := filepath.Base(filePath)
		name := strings.TrimSuffix(fileName, ".md")

		note := &Note{
			Name: name,
			Path: filePath,
		}

		notes = append(notes, note)
	}

	return notes, nil
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
