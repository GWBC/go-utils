package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMd5(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
