package node

import (
	"context"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"log"
	"sync"
)

type ChordNode struct {
	id      []byte
	address string
}

type ChordServer struct {
	self        *ChordNode
	predecessor *ChordNode
	successor   *ChordNode
	known       []*ChordNode
	mux         *sync.Mutex
}

func NewChordServer(addr string) *ChordServer {
	self := &ChordNode{
		generate_sha1(addr),
		addr,
	}
	return &ChordServer{
		self,
		nil,
		self,
		make([]*ChordNode, 0),
		&sync.Mutex{},
	}
}

func (s *ChordServer) FindSuccessor(ctx context.Context, in *pb.FindSuccessorRequest) (*pb.FindSuccessorReply, error) {
	id := in.Id
	if in_range(id, s.self.id, s.successor.id) {
		return &pb.FindSuccessorReply{Id: s.successor.id, Addr: s.successor.address}, nil
	} else {
		node, err := s.successor.FindSuccessor(ctx, in.Id)
		if err != nil {
			return nil, err
		}
		return &pb.FindSuccessorReply{Id: node.id, Addr: node.address}, nil
	}
}

func (s *ChordServer) Join(ctx context.Context, node *ChordNode) error {
	s.mux.Lock()
	s.predecessor = nil
	s.mux.Unlock()

	node, err := node.FindSuccessor(ctx, s.self.id)
	if err != nil {
		return err
	}

	s.mux.Lock()
	s.successor = node
	s.mux.Unlock()
	return nil
}

func (s *ChordServer) Stabilize(ctx context.Context) error {
	return nil
}

func (s *ChordServer) Notify(ctx context.Context, in *pb.NotifyRequest) (*pb.Result, error) {
	return nil, nil
}

func (s *ChordServer) FixFingers(ctx context.Context) error {
	return nil
}

func (s *ChordServer) CheckPredecessor(ctx context.Context) error {
	return nil
}

func (n *ChordNode) FindSuccessor(ctx context.Context, id []byte) (*ChordNode, error) {
	conn, err := grpc.Dial(n.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	r, err := c.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: id})
	if err != nil {
		log.Fatalf("could not request: %v", err)
		return nil, err
	}
	return &ChordNode{r.Id, r.Addr}, nil
}
