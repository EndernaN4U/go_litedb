package litedb

import "fmt"

type LiteDb struct {
	dir_path string
}

func (db *LiteDb) open() {
	fmt.Println("opened")
}
