package litedb

import (
	"path/filepath"
)

type LiteDb struct {
	dir_path string
}

// Create new LiteDb instance
func New(dir_path string) LiteDb {
	return LiteDb{dir_path: dir_path}
}

// Creates new instance of table object
func (db *LiteDb) Table(table_name string) table {
	table_path := filepath.Join(db.dir_path, table_name)

	return table{path: table_path}
}
