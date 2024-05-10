package litedb

import (
	"path/filepath"
)

type LiteDb struct {
	dir_path string
}

func New(dir_path string) LiteDb {
	return LiteDb{dir_path: dir_path}
}

// For less complexity, directory in db should be done before
func (db *LiteDb) Table(table_name string) table {
	table_path := filepath.Join(db.dir_path, table_name)

	return table{path: table_path}
}
