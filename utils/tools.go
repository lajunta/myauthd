package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"strings"
)

var (
	method string
)

// Crypt encrypt string
func Crypt(password string) string {
	var h hash.Hash
	if CFG.CryptMethod == "" {
		return password
	}
	switch CFG.CryptMethod {
	case "sha256":
		h = sha256.New()
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	}
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
