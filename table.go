package litedb

import (
	"os"
	"path/filepath"
)

type table struct {
	path string
}

func (tab *table) NewDoc(data []byte) string {
	uuid := genRandomString(16)
	SaveFile(tab.DocPath(BrackUUID(uuid)), data)

	return uuid
}

func (tab *table) UpdateDoc(id string, data []byte) {
	SaveFile(tab.DocPath(id), data)
}

func (tab *table) DeleteDoc(id string) {
	os.Remove(filepath.Join(tab.path, id))
}

func (tab *table) Doc(id string) []byte {
	data, err := OpenFile(tab.DocPath(id))
	if err != nil {
		return nil
	}
	return data
}

func (tab *table) DocPath(id string) string {
	return filepath.Join(tab.path, id)
}

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
