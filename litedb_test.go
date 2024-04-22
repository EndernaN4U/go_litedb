package litedb

import (
	"encoding/json"
	"testing"
)

func TestMain(t *testing.T) {
	db := LiteDb{dir_path: "./test_db"}
	data, err := db.Open("test.json")

	var buffor map[string]interface{}
	json.Unmarshal(data, &buffor)

	if err != nil {
		t.Error("Error reading")
	}

	if buffor["data"] != "data" {
		t.Errorf("Error")
	}
}

type res struct {
	Data []string `json:"data"`
}

func TestArray(t *testing.T) {
	db := LiteDb{dir_path: "./test_db"}
	data, err := db.Open("array.json")

	buffor := res{}
	json.Unmarshal(data, &buffor)

	if err != nil {
		t.Error("Error reading")
	}

	if buffor.Data[0] != "gg" {
		t.Errorf("Data is wrong. \n Data: %s", buffor)
	}

}
