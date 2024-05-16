package litedb

import (
	"fmt"
	"strconv"
)

func DoID(unid string) int {
	unid = unid[1 : len(unid)-1]
	id, _ := strconv.Atoi(unid)
	return id
}

func UnID(id int) string {
	return fmt.Sprintf("[%d]", id)
}
