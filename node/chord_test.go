package node

import (
	"bytes"
	"context"
	"fmt"
	. "github.com/franela/goblin"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"sort"
	"testing"
	"time"
)

var test_addr = "abcdeabcdeabcdeabcde"
var test_id = []byte(test_addr)

func TestChordNode(t *testing.T) {
	g := Goblin(t)
	g.Describe("bootstrap", func() {
		server := NewChordServer("127.0.0.1:23333")
		g.It("should initialize node", func() {
			g.Assert(len(server.self.id)).Eql(M_bytes)
			g.Assert(server.predecessor == nil).Eql(true)
			g.Assert(server.successor()).Eql(server.self)
		})
	})
}

func TestChordRPC(t *testing.T) {
	g := Goblin(t)
	g.Describe("find successor rpc", func() {
		server := NewChordServer("127.0.0.1:23333")
		g.It("should be self single node network", func() {
			node, err := server.FindSuccessor(context.Background(), &pb.FindSuccessorRequest{Id: test_id})
			g.Assert(err == nil).IsTrue()
			g.Assert(node.Id).Equal(server.self.id)
		})
		g.It("should return given predecessor when start up", func() {
			node, err := server.FindPredecessor(context.Background(), &pb.Node{Id: test_id, Addr: test_addr})
			g.Assert(err == nil).IsTrue()
			g.Assert(node.Addr).Equal(test_addr)
		})
	})
}

func run_server(s *grpc.Server, server *ChordServer, done chan bool) {
	lis, err := net.Listen("tcp", server.self.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pb.RegisterChordServer(s, server)
	pb.RegisterDHTServer(s, server)
	done <- true
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func MakeChordCluster(n int) ([]*grpc.Server, []*ChordServer) {
	log.Printf("making clusters of %d nodes", n)
	done := make(chan bool)
	var chord_servers []*ChordServer
	var grpc_servers [] *grpc.Server
	var addr []string
	addr_base := rand.Intn(1000) + 50000
	for i := 0; i < n; i++ {
		addr = append(addr, fmt.Sprintf("127.0.0.1:%v", i+addr_base))
	}
	sort.SliceStable(addr, func(i, j int) bool {
		return bytes.Compare(generate_sha1(addr[i]), generate_sha1(addr[j])) < 0
	})
	for i := 0; i < n; i++ {
		chord_servers = append(chord_servers, NewChordServer(addr[i]))
		grpc_servers = append(grpc_servers, grpc.NewServer())
		go run_server(grpc_servers[i], chord_servers[i], done)
	}
	for i := 0; i < n; i++ {
		<-done
	}
	log.Printf("... done")
	return grpc_servers, chord_servers
}

func TeardownChordCluster(grpc_servers []*grpc.Server, chord_servers []*ChordServer) {
	log.Printf("tearing down")
	for i := range grpc_servers {
		grpc_servers[i].Stop()
	}
	log.Printf("... done")
}

func chord_system_test_join(g *G, n int) {
	grpc_servers, chord_servers := MakeChordCluster(n)
	ctx, cancel := context.WithCancel(context.Background())
	for i := range chord_servers {
		if i != 0 {
			err := chord_servers[i].Join(ctx, chord_servers[0].self)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}
	}
	for i := range chord_servers {
		g.Assert(chord_servers[i].successor() == nil).IsFalse()
		id1 := chord_servers[i].successor().id
		id2 := chord_servers[i].self.id
		if i != 0 {
			g.Assert(bytes.Equal(id1, id2)).IsFalse()
		}
	}
	cancel()
	TeardownChordCluster(grpc_servers, chord_servers)
}

func chord_system_test_stabilization(g *G, n int) {
	grpc_servers, chord_servers := MakeChordCluster(n)
	ctx, cancel := context.WithCancel(context.Background())
	for i := range chord_servers {
		if i != 0 {
			err := chord_servers[i].Join(ctx, chord_servers[0].self)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := range chord_servers {
			err := chord_servers[i].Stabilize(ctx)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}
	}
	for i := range chord_servers {
		next_i := (i + 1) % len(chord_servers)
		g.Assert(chord_servers[i].successor() == nil).IsFalse()
		g.Assert(chord_servers[i].predecessor == nil).IsFalse()
		g.Assert(chord_servers[next_i].successor() == nil).IsFalse()
		g.Assert(chord_servers[next_i].predecessor == nil).IsFalse()
		g.Assert(bytes.Equal(chord_servers[next_i].predecessor.id, chord_servers[i].self.id)).IsTrue()
		g.Assert(bytes.Equal(chord_servers[i].successor().id, chord_servers[next_i].self.id)).IsTrue()
	}
	cancel()
	TeardownChordCluster(grpc_servers, chord_servers)
}

func TestChordSystem(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping system testing in short mode")
	}

	g := Goblin(t)

	g.Describe("node join", func() {
		g.It("should join a 10-node network", func() {
			chord_system_test_join(g, 10)
		})
		g.It("should join a 50-node network", func() {
			chord_system_test_join(g, 50)
		})
	})

	g.Describe("node stabilization", func() {
		g.It("should stabilize a 10-node network", func() {
			chord_system_test_stabilization(g, 10)
		})
		g.It("should stabilize a 50-node network", func() {
			g.Timeout(time.Second * 30)
			chord_system_test_stabilization(g, 50)
		})
	})
}
