package node

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestUtils(t *testing.T) {
	g := Goblin(t)
	g.Describe("Chord Utils", func() {
		g.It("should generate 160-bit hash", func() {
			s := "127.0.0.1:23333"
			h := generate_sha1(s)
			g.Assert(len(h)).Equal(160 / 8)
		})
	})
}
