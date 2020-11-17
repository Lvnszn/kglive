package util

import (
	"crypto/md5"
	"fmt"
)

// MD5Encode encode string to hex
func MD5Encode(in string) (out string) {
	data := []byte(in)
	has := md5.Sum(data)
	out = fmt.Sprintf("%x", has)
	return
}
