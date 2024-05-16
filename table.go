package litedb

import (
	"os"
	"path/filepath"
)

type table struct {
	path string
}

func (tab *table) NewDoc(data []byte) int {
	cache := tab.cacheData()

	cache.Last_Index++
	SaveFile(tab.DocPath(cache.Last_Index), data)

	tab.cacheUpdate(cache)

	return cache.Last_Index
}

func (tab *table) UpdateDoc(id int, data []byte) {
	SaveFile(tab.DocPath(id), data)
}

func (tab *table) DeleteDoc(id int) {
	sid := UnID(id)

	os.Remove(filepath.Join(tab.path, sid))
}

func (tab *table) Doc(id int) []byte {
	data, err := OpenFile(tab.DocPath(id))
	if err != nil {
		return nil
	}
	return data
}

func (tab *table) DocPath(id int) string {
	return filepath.Join(tab.path, UnID(id))
}

func (tab *table) IDs() []string {
	// TODO: all tab ids with deleted ".cache"
	return []string{}
}

func (tab *table) LastIndex() int {
	return tab.cacheData().Last_Index
}
