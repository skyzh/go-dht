package node

import (
	"bytes"
	"context"
	"errors"
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
		nil,
		&sync.Mutex{},
	}
}

func (s *ChordServer) FindSuccessor(ctx context.Context, in *pb.FindSuccessorRequest) (*pb.Node, error) {
	id := in.Id
	s.mux.Lock()
	if in_range(id, s.self.id, s.successor.id) {
		defer s.mux.Unlock()
		return &pb.Node{Id: s.successor.id, Addr: s.successor.address}, nil
	} else {
		id := in.Id
		s.mux.Unlock()
		node, err := s.successor.FindSuccessor(ctx, id)
		if err != nil {
			return nil, err
		}
		s.mux.Lock()
		defer s.mux.Unlock()
		return &pb.Node{Id: node.id, Addr: node.address}, nil
	}
}

func (s *ChordServer) Join(ctx context.Context, node *ChordNode) error {
	s.mux.Lock()
	s.predecessor = nil
	id := s.self.id
	s.mux.Unlock()

	node, err := node.FindSuccessor(ctx, id)
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
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.predecessor == nil ||
		in_range(in.Id, s.predecessor.id, s.self.id) && !bytes.Equal(in.Id, s.self.id) {
		s.predecessor = &ChordNode{in.Id, in.Addr}
	}
	return &pb.Result{Result: "success"}, nil
}

func (s *ChordServer) FindPredecessor(ctx context.Context, in *pb.Void) (*pb.Node, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.predecessor == nil {
		return nil, errors.New("no predecessor")
	}
	return &pb.Node{Id: s.predecessor.id, Addr: s.predecessor.address}, nil
}

func (s *ChordServer) FixFingers(ctx context.Context) error {
	return nil
}

func (s *ChordServer) CheckPredecessor(ctx context.Context) error {
	return nil
}

func (s *ChordServer) ClosestPrecedingNode(ctx context.Context, id []byte) (*ChordNode, error) {
	return nil, nil
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
