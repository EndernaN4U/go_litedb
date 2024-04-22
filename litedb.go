package litedb

import (
	"os"
	"path/filepath"
)

type LiteDb struct {
	dir_path string
}

func (db *LiteDb) Open(file_name string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(db.dir_path, file_name))
	return data, err
}

func (db *LiteDb) Save(file_name string, data []byte) {
	os.WriteFile(
		filepath.Join(db.dir_path, file_name),
		data,
		0644,
	)
}
