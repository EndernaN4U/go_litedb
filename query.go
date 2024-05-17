package litedb

// Numeric values should be float
type Criterion struct {
	Name  string
	Value any
}

func (tab *table) Filter(criteria []Criterion) []string {
	res := []string{}

	for _, id := range tab.IDs() {
		data, _ := OpenFile(tab.DocPath(id))
		data_json := BytesToJson[map[string]interface{}](data)

		flag := true
		for _, criterion := range criteria {
			if data_json[criterion.Name] != criterion.Value {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, id)
		}
	}

	return res
}
