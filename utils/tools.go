package utils

import (
	"crypto/sha256"
	"strings"
)

// Crypt encrypt string
func Crypt(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	return string(bs)
}

// Filter bypass some sql attack
func Filter(str string) string {
	strs := []string{" ", "\\", "/", "&", ";"}
	for _, v := range strs {
		str = strings.ReplaceAll(str, v, "")
	}
	return str
}
