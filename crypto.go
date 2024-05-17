package litedb

import (
	"crypto/rand"
	"encoding/base64"
)

func genRandomBytes(size int) []byte {
	buf := make([]byte, size)
	rand.Read(buf)

	return buf
}

func genRandomString(bytes_am int) string {
	return base64.URLEncoding.EncodeToString(genRandomBytes(bytes_am))
}
