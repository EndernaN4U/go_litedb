package litedb

import (
	"fmt"
)

// Add brackets around uuid (<uuid> => [<uuid>] )
func BrackUUID(uuid string) string {
	return fmt.Sprintf("[%s]", uuid)
}

// Removes brackets around [uuid] ( [<uuid>] => <uuid> )
func UnbrackUUID(uuid string) string {
	return uuid[1 : len(uuid)-1]
}
