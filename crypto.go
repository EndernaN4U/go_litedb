package litedb

import (
	"crypto/rand"
	"encoding/base64"
)

// Generate and return random bytes array
func genRandomBytes(size int) []byte {
	buf := make([]byte, size)
	rand.Read(buf)

	return buf
}

// Generate and return random string using "genRandomBytes"
func genRandomString(bytes_am int) string {
	return base64.URLEncoding.EncodeToString(genRandomBytes(bytes_am))
}
