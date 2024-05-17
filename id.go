package litedb

import (
	"fmt"
)

func BrackUUID(uuid string) string {
	return fmt.Sprintf("[%s]", uuid)
}

func UnbrackUUID(uuid string) string {
	return uuid[1 : len(uuid)-1]
}
