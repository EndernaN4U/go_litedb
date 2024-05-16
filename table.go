package litedb

import (
	"fmt"
	"os"
	"path/filepath"
)

type table struct {
	path string
}

func (tab *table) NewDoc(data []byte) int {
	cache := tab.cacheData()

	cache.Last_Index++
	id := fmt.Sprintf("[%d]", cache.Last_Index)

	tab.cacheUpdate(cache)

	SaveFile(filepath.Join(tab.path, id), data)

	return cache.Last_Index
}

func (tab *table) UpdateDoc(id int, data []byte) {
	sid := fmt.Sprintf("[%d]", id)

	SaveFile(filepath.Join(tab.path, sid), data)
}

func (tab *table) DeleteDoc(id int) {
	sid := UnID(id)

	os.Remove(filepath.Join(tab.path, sid))
}

func (tab *table) Doc(id int) []byte {
	data, err := OpenFile(filepath.Join(tab.path, fmt.Sprintf("[%d]", id)))
	if err != nil {
		return nil
	}
	return data
}

// TODO: add tab doc path

func (tab *table) IDs() []string {
	// TODO: all tab ids with deleted ".cache"
	return []string{}
}

func (tab *table) LastIndex() int {
	return tab.cacheData().Last_Index
}
