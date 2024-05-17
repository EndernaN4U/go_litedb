package litedb

import (
	"os"
	"path/filepath"
)

type table struct {
	path string
}

// Creates new document and returns its [uuid]
func (tab *table) NewDoc(data []byte) string {
	uuid := BrackUUID(genRandomString(16))
	SaveFile(tab.DocPath(uuid), data)

	return uuid
}

// Returns document value in bytes
func (tab *table) Doc(id string) []byte {
	data, _ := OpenFile(tab.DocPath(id))
	return data
}

// Updates document using same values
func (tab *table) UpdateDoc(id string, data []byte) {
	SaveFile(tab.DocPath(id), data)
}

// MAYBE: Updating could be using criteria

// Deletes document
func (tab *table) DeleteDoc(id string) {
	os.Remove(filepath.Join(tab.path, id))
}

// Returns absolute path of document
func (tab *table) DocPath(id string) string {
	return filepath.Join(tab.path, id)
}

// Returns all [uuid]s that table contains
func (tab *table) IDs() []string {
	files, _ := os.ReadDir(tab.path)

	files_names := []string{}
	for _, file := range files {
		file_name := file.Name()

		if file_name == ".cache" {
			continue
		}

		files_names = append(files_names, file_name)
	}

	return files_names
}
