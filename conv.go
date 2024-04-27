package litedb

import "encoding/json"

func BytesToJson[T any](data []byte) T {
	var json_data T
	json.Unmarshal(data, &json_data)

	return json_data
}
