package litedb

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type table struct {
	path string
}

func (tab *table) NewDoc(data []byte) int {
	cache := cacheData(*tab)

	cache.Last_Index++
	id := fmt.Sprintf("[%d]", cache.Last_Index)

	cacheUpdate(*tab, cache)

	SaveFile(filepath.Join(tab.path, id), data)

	return cache.Last_Index
}

func (tab *table) UpdateDoc(id int, data []byte) {
	sid := fmt.Sprintf("[%d]", id)

	SaveFile(filepath.Join(tab.path, sid), data)
}

func (tab *table) DeleteDoc(id int) {
	sid := fmt.Sprintf("[%d]", id)

	os.Remove(filepath.Join(tab.path, sid))
}

func (tab *table) Doc(id int) []byte {
	data, err := OpenFile(filepath.Join(tab.path, fmt.Sprintf("[%d]", id)))
	if err != nil {
		return nil
	}
	return data
}

func (tab *table) Docs() []int {
	dirs := ReadDir(tab.path)
	ids := []int{}

	for _, dir := range dirs {
		dir = dir[1 : len(dir)-1]
		dir_ind, _ := strconv.Atoi(dir)
		ids = append(ids, dir_ind)
	}

	return ids
}

func (tab *table) LastIndex() int {
	cache := cacheData(*tab)
	return cache.Last_Index
}
