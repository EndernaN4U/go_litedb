package litedb

import "encoding/json"

// Converts bytes array to object instance using generic
func BytesToJson[T any](data []byte) T {
	var json_data T
	json.Unmarshal(data, &json_data)

	return json_data
}
