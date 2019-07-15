package node

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestUtils(t *testing.T) {
	g := Goblin(t)
	g.Describe("generate_sha1", func() {
		g.It("should generate 160-bit hash", func() {
			s := "127.0.0.1:23333"
			h := generate_sha1(s)
			g.Assert(len(h)).Equal(160 / 8)
		})
	})
	g.Describe("in_range", func() {
		g.It("should not include left point", func() {
			a := []byte("11111111111111111111")
			c := a
			b := []byte("22222222222222222222")
			g.Assert(in_range(c, a, b)).IsFalse()
		})
		g.It("should include right point", func() {
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
		})
	})
}
