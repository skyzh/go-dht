package node

import (
	"crypto/sha1"
)

func generate_sha1(s string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	bs := hasher.Sum(nil)
	return bs
}
