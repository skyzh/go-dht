package node

import (
	"bytes"
	. "github.com/franela/goblin"
	"math/big"
	"testing"
)

func TestUtils(t *testing.T) {
	g := Goblin(t)
	g.Describe("generate_chord_hash", func() {
		g.It("should generate M-bit hash", func() {
			s := "127.0.0.1:23333"
			h := generate_chord_hash(s)
			g.Assert(len(h)).Equal(M_bytes)
		})
	})
	g.Describe("generate_kad_hash", func() {
		g.It("should generate 160-bit hash", func() {
			s := "127.0.0.1:23333"
			h := generate_kad_hash(s)
			g.Assert(len(h)).Equal(160 / 8)
		})
	})
	g.Describe("xor_distance", func() {
		g.It("should calculate distance", func() {
			a := make([]byte, 20)
			b := []byte("23333233332333323333")
			g.Assert(string(xor_distance(a, b))).Equal(string(b))
			a = []byte("23333233332333323333")
			b = []byte("23333233332333323333")
			g.Assert(bytes.Equal(xor_distance(a, b), []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})).IsTrue()
		})
	})
	g.Describe("in_range", func() {
		g.It("should not include left bound", func() {
			a := []byte("11111111111111111111")
			c := a
			b := []byte("22222222222222222222")
			g.Assert(in_range(c, a, b)).IsFalse()
		})
		g.It("should include right bound", func() {
			a := []byte("11111111111111111111")
			b := []byte("22222222222222222222")
			c := b
			g.Assert(in_range(c, a, b)).IsTrue()
		})
		g.It("should process cycle", func() {
			b := []byte("11111111111111111111")
			a := []byte("22222222222222222222")
			c := []byte("33333333333333333333")
			g.Assert(in_range(c, a, b)).IsTrue()
		})
		g.It("should include self in cycle", func() {
			b := []byte("11111111111111111111")
			a := []byte("11111111111111111111")
			c := []byte("11111111111111111111")
			g.Assert(in_range(c, a, b)).IsTrue()
			d := []byte("11111111211111111111")
			g.Assert(in_range(d, a, b)).IsTrue()
		})
	})
	g.Describe("in_range_exclude", func() {
		g.It("should not include left bound", func() {
			a := []byte("11111111111111111111")
			c := a
			b := []byte("22222222222222222222")
			g.Assert(in_range_exclude(c, a, b)).IsFalse()
		})
		g.It("should not include right bound", func() {
			a := []byte("11111111111111111111")
			b := []byte("22222222222222222222")
			c := b
			g.Assert(in_range_exclude(c, a, b)).IsFalse()
		})
		g.It("should process cycle", func() {
			b := []byte("11111111111111111111")
			a := []byte("22222222222222222222")
			c := []byte("33333333333333333333")
			g.Assert(in_range_exclude(c, a, b)).IsTrue()
		})
		g.It("should include anything in cycle", func() {
			b := []byte("11111111111111111111")
			a := []byte("11111111111111111111")
			d := []byte("11111111211111111111")
			g.Assert(in_range(d, a, b)).IsTrue()
		})
	})
	g.Describe("byte_add_power_2", func() {
		g.It("should add correctly", func() {
			a := big.NewInt(233)
			b := big.NewInt(0)
			b.SetBytes(byte_add_power_2(a.Bytes(), 10))
			g.Assert(b.Cmp(big.NewInt(233+1024)) == 0).IsTrue()
			b.SetBytes(byte_add_power_2(a.Bytes(), 16))
			if M > 16 {
				g.Assert(b.Cmp(big.NewInt(233+65536)) == 0).IsTrue()
			}
		})
		g.It("should add beyond bound", func() {
			a := big.NewInt(233)
			b := big.NewInt(0)
			b.SetBytes(byte_add_power_2(a.Bytes(), M+2))
			g.Assert(b.Cmp(big.NewInt(233)) == 0).IsTrue()
			a.Lsh(big.NewInt(1), M-1)
			b.SetBytes(byte_add_power_2(a.Bytes(), M-1))
			g.Assert(b.Cmp(big.NewInt(0)) == 0).IsTrue()
		})
	})
}
