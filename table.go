package litedb

import (
	"fmt"
	"path/filepath"
)

type table struct {
	path string
}

func (tab *table) NewDoc(data []byte) int {
	id := "[0]"
	SaveFile(filepath.Join(tab.path, id), data)
	return 0
}

func (tab *table) UpdateDoc(id int, data []byte) {
	sid := fmt.Sprintf("[%d]", id)
	SaveFile(filepath.Join(tab.path, sid), data)
}

func (tab *table) Doc(id int) []byte {
	data, err := OpenFile(filepath.Join(tab.path, fmt.Sprintf("[%d]", id)))
	if err != nil {
		return nil
	}
	return data
}
