package crypt

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
)

// Scrypt 字符串加密
func Scrypt(s string) string {
	const KeyLen = 32
	salt := []byte{6, 10, 11, 16, 23, 25, 13, 66}
	hash, err := scrypt.Key([]byte(s), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	ret := base64.StdEncoding.EncodeToString(hash)
	return ret
}
