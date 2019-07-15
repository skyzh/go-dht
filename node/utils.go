package node

import (
	"bytes"
	"crypto/sha1"
)

func generate_sha1(s string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	bs := hasher.Sum(nil)
	return bs
}

func in_range(c, l, r []byte) bool {
	if bytes.Compare(l, r) < 0 {
		return bytes.Compare(l, c) < 0 && bytes.Compare(c, r) <= 0
	} else {
		return bytes.Compare(l, c) < 0 || bytes.Compare(c, r) <= 0
	}
	return true
}
