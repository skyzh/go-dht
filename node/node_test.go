package node

import (
	"context"
	. "github.com/franela/goblin"
	pb "github.com/skyzh/go-dht/chord"
	"testing"
)

func TestNode(t *testing.T) {
	g := Goblin(t)
	g.Describe("Chord RPCs", func() {
		server := &DHTServer{}
		g.It("should return successor ", func() {
			request := &pb.SuccessorRequest{}
			reply, err := server.FindSuccessor(context.Background(), request)
			g.Assert(reply.Id).Equal("23333")
			g.Assert(err).Equal(nil)
		})
	})
}
