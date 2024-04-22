package litedb

import (
	"path/filepath"
)

type LiteDb struct {
	dir_path string
}

// For less complexity folder in db should be done before
func (db *LiteDb) Table(table_name string) table {
	return table{path: filepath.Join(db.dir_path, table_name)}
}
