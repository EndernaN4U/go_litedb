package litedb

import (
	"os"
)

type LiteDb struct {
	dir_path string
}

func (db *LiteDb) open() (string, error) {
	data, err := os.ReadFile(db.dir_path)
	return string(data), err
}
