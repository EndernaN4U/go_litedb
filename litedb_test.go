package litedb

import "testing"

func TestMain(t *testing.T) {
	db := LiteDb{dir_path: "./dbs"}
	db.open()
}
