package litedb

import (
	"path/filepath"
)

// Numeric values should be float
type Criterion struct {
	Name  string
	Value any
}

func (tab *table) Filter(criteria []Criterion) []int {
	res := []int{}

	for _, id := range tab.IDs() {
		data, _ := OpenFile(filepath.Join(tab.path, id))
		data_json := BytesToJson[map[string]interface{}](data)

		flag := true
		for _, criterion := range criteria {
			if data_json[criterion.Name] != criterion.Value {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, DoID(id))
		}
	}

	return res
}
