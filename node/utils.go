package node

import (
	"bytes"
	"crypto/sha1"
	"math/big"
)

func generate_chord_hash(s string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	bs := hasher.Sum(nil)
	return bs[:M_bytes]
}

func generate_kad_hash(s string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return hasher.Sum(nil)
}

func xor_distance(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c
}

func in_range(c, l, r []byte) bool {
	if bytes.Compare(l, r) < 0 {
		return bytes.Compare(l, c) < 0 && bytes.Compare(c, r) <= 0
	} else {
		return bytes.Compare(l, c) < 0 || bytes.Compare(c, r) <= 0
	}
}

func in_range_exclude(c, l, r []byte) bool {
	return !bytes.Equal(c, r) && in_range(c, l, r)
}

func byte_add_power_2(id []byte, exp uint) []byte {
	z := new(big.Int)
	z.SetBytes(id)
	p := new(big.Int)
	p.Lsh(big.NewInt(1), exp)
	z.Add(z, p)
	b := z.Bytes()
	result := make([]byte, M_bytes)
	if len(b) > M_bytes {
		for i := 0; i < M_bytes; i++ {
			result[i] = b[i+len(b)-M_bytes]
		}
	} else {
		i := 0
		for ; i < M_bytes-len(b); i++ {
			result[i] = 0
		}
		for ; i < M_bytes; i++ {
			result[i] = b[i+len(b)-M_bytes]
		}
	}
	return result
}
