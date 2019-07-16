package node

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"sync"
)

type KadNode struct {
	Id      []byte
	Address string
}

type KadServer struct {
	self    *KadNode
	mux     *sync.Mutex
	logger  *log.Entry
	storage map[string]string
}

func NewKadServer(addr string) *KadServer {
	self := &KadNode{
		generate_kad_hash(addr),
		addr,
	}
	return &KadServer{
		self,
		&sync.Mutex{},
		log.WithFields(log.Fields{
			"id": fmt.Sprintf("%X", self.Id),
		}),
		make(map[string]string),
	}
}

func (s *KadServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{Id: s.self.Id, Addr: s.self.Address}, nil
}

func (s *KadServer) FindNode(ctx context.Context, in *pb.FindNodeRequest) (*pb.FindNodeReply, error) {
	return &pb.FindNodeReply{Id: s.self.Id, Addr: s.self.Address}, nil
}

func (s *KadServer) FindValue(ctx context.Context, in *pb.FindValueRequest) (*pb.FindValueReply, error) {
	return &pb.FindValueReply{Id: s.self.Id, Addr: s.self.Address}, nil
}

func (s *KadServer) Store(ctx context.Context, in *pb.StoreRequest) (*pb.StoreReply, error) {
	return &pb.StoreReply{Id: s.self.Id, Addr: s.self.Address}, nil
}

func (n *KadNode) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewKadClient(conn)
	r, err := c.Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (n *KadNode) FindNode(ctx context.Context, in *pb.FindNodeRequest) (*pb.FindNodeReply, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewKadClient(conn)
	r, err := c.FindNode(ctx, in)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (n *KadNode) FindValue(ctx context.Context, in *pb.FindValueRequest) (*pb.FindValueReply, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewKadClient(conn)
	r, err := c.FindValue(ctx, in)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (n *KadNode) Store(ctx context.Context, in *pb.StoreRequest) (*pb.StoreReply, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewKadClient(conn)
	r, err := c.Store(ctx, in)
	if err != nil {
		return nil, err
	}
	return r, nil
}
