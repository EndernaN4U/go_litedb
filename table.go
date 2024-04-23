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
	new_id := tab.MaxIndex() + 1
	id := fmt.Sprintf("[%d]", new_id)

	SaveFile(filepath.Join(tab.path, id), data)

	return new_id
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

func (tab *table) MaxIndex() int {
	dirs := ReadDir(tab.path)
	max_ind := -1

	for _, dir := range dirs {
		dir = dir[1 : len(dir)-1]
		dir_ind, _ := strconv.Atoi(dir)
		if dir_ind > max_ind {
			max_ind = dir_ind
		}
	}

	return max_ind
}
