package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"strings"

	"golang.org/x/text/encoding/charmap"
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

// Latin1ToUtf8 convert latin1 string to utf8
func Latin1ToUtf8(s string) string {
	b := []byte{}
	buf := bytes.NewBufferString(s)
	r := charmap.ISO8859_1.NewDecoder().Reader(buf)
	r.Read(b)
	return string(b)
}
