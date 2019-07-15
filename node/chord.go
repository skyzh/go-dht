package node

import (
	"context"
	"errors"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"sync"
)

const (
	M       = 160
	M_bytes = M / 8
)

type ChordNode struct {
	id      []byte
	address string
}

type ChordServer struct {
	self            *ChordNode
	predecessor     *ChordNode
	finger          []*ChordNode
	fix_finger_next uint
	mux             *sync.Mutex
}

func (s *ChordServer) successor() (*ChordNode) {
	if len(s.finger) >= 1 {
		return s.finger[0]
	} else {
		return nil
	}
}

func NewChordServer(addr string) *ChordServer {
	self := &ChordNode{
		generate_sha1(addr),
		addr,
	}
	return &ChordServer{
		self,
		nil,
		[]*ChordNode{self},
		0,
		&sync.Mutex{},
	}
}

func (s *ChordServer) FindSuccessor(ctx context.Context, in *pb.FindSuccessorRequest) (*pb.Node, error) {
	s.mux.Lock()
	if in_range(in.Id, s.self.id, s.successor().id) {
		defer s.mux.Unlock()
		return &pb.Node{Id: s.successor().id, Addr: s.successor().address}, nil
	} else {
		s.mux.Unlock()
		n, err := s.ClosestPrecedingNode(ctx, in.Id)
		if err != nil {
			return nil, err
		}
		node, err := n.FindSuccessor(ctx, in.Id)
		if err != nil {
			return nil, err
		}
		return &pb.Node{Id: node.id, Addr: node.address}, nil
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
	s.finger[0] = node
	s.mux.Unlock()
	return nil
}

func (s *ChordServer) Stabilize(ctx context.Context) error {
	x, err := s.successor().FindPredecessor(ctx, s.self)
	if err != nil {
		return err
	}
	if in_range_exclude(x.id, s.self.id, s.successor().id) {
		s.finger[0] = x
	}
	err = s.successor().Notify(ctx, s.self)
	if err != nil {
		return err
	}
	return nil
}

func (s *ChordServer) Notify(ctx context.Context, in *pb.Node) (*pb.Result, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.predecessor == nil ||
		in_range_exclude(in.Id, s.predecessor.id, s.self.id) {
		s.predecessor = &ChordNode{in.Id, in.Addr}
	}
	return &pb.Result{Result: "success"}, nil
}

func (s *ChordServer) FindPredecessor(ctx context.Context, in *pb.Node) (*pb.Node, error) {
	_, err := s.Notify(ctx, in)
	if err != nil {
		return nil, err
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.predecessor == nil {
		return nil, errors.New("no predecessor")
	}
	return &pb.Node{Id: s.predecessor.id, Addr: s.predecessor.address}, nil
}

func (s *ChordServer) Ping(ctx context.Context, in *pb.Node) (*pb.Void, error) {
	return &pb.Void{}, nil
}

func (s *ChordServer) FixFingers(ctx context.Context) error {
	s.mux.Lock()
	s.fix_finger_next = s.fix_finger_next + 1
	if s.fix_finger_next >= M {
		s.fix_finger_next = 0
	}
	s.mux.Unlock()
	x, err := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: byte_add_power_2(s.self.id, s.fix_finger_next)})
	if err != nil {
		return err
	}
	s.mux.Lock()
	s.finger[s.fix_finger_next] = &ChordNode{x.Id, x.Addr}
	defer s.mux.Unlock()
	return nil
}

func (s *ChordServer) CheckPredecessor(ctx context.Context) error {
	return nil
}

func (s *ChordServer) ClosestPrecedingNode(ctx context.Context, id []byte) (*ChordNode, error) {
	for i := M - 1; i >= 0; i-- {
		if in_range_exclude(s.finger[i].id, s.self.id, id) {
			return s.finger[i], nil
		}
	}
	return nil, nil
}

func (n *ChordNode) FindSuccessor(ctx context.Context, id []byte) (*ChordNode, error) {
	conn, err := grpc.Dial(n.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	r, err := c.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &ChordNode{r.Id, r.Addr}, nil
}

func (n *ChordNode) FindPredecessor(ctx context.Context, self *ChordNode) (*ChordNode, error) {
	conn, err := grpc.Dial(n.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	r, err := c.FindPredecessor(ctx, &pb.Node{Id: self.id, Addr: self.address})
	if err != nil {
		return nil, err
	}
	return &ChordNode{r.Id, r.Addr}, nil
}

func (n *ChordNode) Notify(ctx context.Context, self *ChordNode) error {
	conn, err := grpc.Dial(n.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	_, err = c.Notify(ctx, &pb.Node{Id: self.id, Addr: self.address})
	if err != nil {
		return err
	}
	return nil
}
