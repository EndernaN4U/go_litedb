package litedb

import (
	"fmt"
	"strconv"
)

func BrackUUID(uuid string) string {
	return fmt.Sprintf("[%s]", uuid)
}

func UnbrackUUID(uuid string) string {
	return uuid[1 : len(uuid)-1]
}

func DoID(unid string) int {
	unid = unid[1 : len(unid)-1]
	id, _ := strconv.Atoi(unid)
	return id
}

func UnID(id int) string {
	return fmt.Sprintf("[%d]", id)
}
