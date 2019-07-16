package node

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"sync"
	"time"
)

const (
	M       = 16
	M_bytes = M / 8
)

type ChordNode struct {
	Id      []byte
	Address string
}

type ChordServer struct {
	self            *ChordNode
	predecessor     *ChordNode
	finger          []*ChordNode
	fix_finger_next uint
	mux             *sync.Mutex
	logger          *log.Entry
	storage         map[string]string
}

func (s *ChordServer) successor() (*ChordNode) {
	if len(s.finger) >= 1 {
		return s.finger[0]
	} else {
		return nil
	}
}

func NewChordNode(addr string) *ChordNode {
	return &ChordNode{
		Id:      generate_hash(addr),
		Address: addr,
	}
}

func NewChordServer(addr string) *ChordServer {
	self := &ChordNode{
		generate_hash(addr),
		addr,
	}
	finger := make([]*ChordNode, M)
	finger[0] = self
	for i := 1; i < M; i++ {
		finger[i] = nil
	}
	return &ChordServer{
		self,
		nil,
		finger,
		0,
		&sync.Mutex{},
		log.WithFields(log.Fields{
			"id": fmt.Sprintf("%X", self.Id),
		}),
		make(map[string]string),
	}
}

func (s *ChordServer) Serve(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			break
		default:
			{
				err := s.Stabilize(ctx)
				if err != nil {
					s.logger.Warningf("%X routine error %v", s.self.Id, err)
				}
				time.Sleep(time.Millisecond * 5)
			}
			{
				err := s.CheckPredecessor(ctx)
				if err != nil {
					s.logger.Warningf("%X routine error %v", s.self.Id, err)
				}
				time.Sleep(time.Millisecond * 5)
			}
			{
				err := s.FixFingers(ctx)
				if err != nil {
					s.logger.Warningf("%X routine error %v", s.self.Id, err)
				}
				time.Sleep(time.Millisecond * 5)
			}
		}
	}

}

func (s *ChordServer) FindSuccessor(ctx context.Context, in *pb.FindSuccessorRequest) (*pb.Node, error) {
	s.mux.Lock()
	if in_range(in.Id, s.self.Id, s.successor().Id) {
		defer s.mux.Unlock()
		return &pb.Node{Id: s.successor().Id, Addr: s.successor().Address}, nil
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
		return &pb.Node{Id: node.Id, Addr: node.Address}, nil
	}
}

func (s *ChordServer) Join(ctx context.Context, node *ChordNode) error {
	s.mux.Lock()
	s.predecessor = nil
	s.mux.Unlock()

	node, err := node.FindSuccessor(ctx, s.self.Id)
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
	if in_range_exclude(x.Id, s.self.Id, s.successor().Id) {
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
		in_range_exclude(in.Id, s.predecessor.Id, s.self.Id) {
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
	return &pb.Node{Id: s.predecessor.Id, Addr: s.predecessor.Address}, nil
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
	next := s.fix_finger_next
	s.mux.Unlock()
	x, err := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: byte_add_power_2(s.self.Id, s.fix_finger_next)})
	if err != nil {
		return err
	}
	s.mux.Lock()
	s.finger[next] = &ChordNode{x.Id, x.Addr}
	defer s.mux.Unlock()
	return nil
}

func (s *ChordServer) CheckPredecessor(ctx context.Context) error {
	_, err := s.Notify(ctx, &pb.Node{Id: s.self.Id, Addr: s.self.Address})
	if err != nil {
		s.mux.Lock()
		defer s.mux.Unlock()
		s.predecessor = nil
	}
	return nil
}

func (s *ChordServer) ClosestPrecedingNode(ctx context.Context, id []byte) (*ChordNode, error) {
	for i := M - 1; i >= 0; i-- {
		if s.finger[i] == nil {
			continue
		}
		if in_range_exclude(s.finger[i].Id, s.self.Id, id) {
			return s.finger[i], nil
		}
	}
	return nil, nil
}

func (n *ChordNode) FindSuccessor(ctx context.Context, id []byte) (*ChordNode, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
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
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	r, err := c.FindPredecessor(ctx, &pb.Node{Id: self.Id, Addr: self.Address})
	if err != nil {
		return nil, err
	}
	return &ChordNode{r.Id, r.Addr}, nil
}

func (n *ChordNode) Notify(ctx context.Context, self *ChordNode) error {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	_, err = c.Notify(ctx, &pb.Node{Id: self.Id, Addr: self.Address})
	if err != nil {
		return err
	}
	return nil
}
