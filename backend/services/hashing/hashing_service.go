package hashing

import (
	"crypto/md5"
	"fmt"
)

func Hash(value string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(value)))
}
