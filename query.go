package litedb

import (
	"fmt"
	"path/filepath"
)

type Criterion struct {
	Name  string
	Value any
}

func (tab *table) Filter(criteria []Criterion) []int {
	res := []int{}

	for _, id := range ReadDir(tab.path) {
		data, _ := OpenFile(filepath.Join(tab.path, id))
		data_json := BytesToJson[map[string]interface{}](data)

		flag := true
		for _, criterion := range criteria {
			fmt.Println(data_json[criterion.Name])
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
