package litedb

import (
	"testing"
)

func TestMain(t *testing.T) {
	db := LiteDb{dir_path: "./test_db/test.txt"}
	data, err := db.open()
	if err != nil {
		t.Error("Error reading")
	}
	if data != "data=data" {
		t.Errorf("Data doesn't match: %s", data)
	}
}
