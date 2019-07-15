package node

import (
	"context"
	. "github.com/franela/goblin"
	pb "github.com/skyzh/go-dht/protos"
	"testing"
)

var test_id = []byte("abcdeabcdeabcdeabcde")

func TestChordNode(t *testing.T) {
	g := Goblin(t)
	g.Describe("bootstrap", func() {
		server := NewChordServer("127.0.0.1:23333")
		g.It("should initialize node", func() {
			g.Assert(len(server.self.id)).Eql(160 / 8)
			g.Assert(server.predecessor == nil).Eql(true)
			g.Assert(server.successor).Eql(server.self)
		})
	})
}

func TestChordRPC(t *testing.T) {
	g := Goblin(t)
	g.Describe("find successor rpc", func() {
		server := NewChordServer("127.0.0.1:23333")
		g.It("should be self single node network", func() {
			node, err := server.FindSuccessor(context.Background(), &pb.FindSuccessorRequest{ Id: test_id })
			g.Assert(err == nil).IsTrue()
			g.Assert(node.Id).Equal(server.self.id)
		})
	})
}
