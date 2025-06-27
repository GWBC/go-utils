package utils

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// //////////////////////////////////////////////////////////////////////

// 填充字符
func FillStr(s string, maxLen int, fillChar rune) string {
	if len(s) > maxLen {
		return s[:maxLen]
	}

	return s + strings.Repeat(string(fillChar), maxLen-len(s))
}

// 随机字符串
func RandomString(length int) string {
	str := make([]byte, length)
	for i := range str {
		str[i] = charset[rand.Intn(len(charset))]
	}

	return string(str)
}
