package node

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

func (s *ChordServer) Get(ctx context.Context, in *pb.Key) (*pb.Pair, error) {
	logger := s.logger.WithFields(log.Fields{"op": "get", "key": in.Key})
	logger.Infof("request")
	hash := generate_hash(in.Key)
	node, err := s.ClosestPrecedingNode(ctx, hash)
	if err != nil {
		return nil, err
	}
	if node == nil {
		logger.Infof("key belongs to %X", s.self.Id)
		val, ok := s.storage[in.Key]
		if !ok {
			return nil, errors.New("key not found")
		} else {
			return &pb.Pair{Key: in.Key, Value: val}, nil
		}
	} else {
		logger.Infof("forwarding request to %X", node.Id)
		return node.Get(ctx, in)
	}
}

func (s *ChordServer) Put(ctx context.Context, in *pb.Pair) (*pb.Result, error) {
	logger := s.logger.WithFields(log.Fields{"op": "put", "key": in.Key})
	logger.Infof("request")
	hash := generate_hash(in.Key)
	node, err := s.ClosestPrecedingNode(ctx, hash)
	if err != nil {
		return nil, err
	}
	if node == nil {
		logger.Infof("key belongs to %X", s.self.Id)
		s.mux.Lock()
		defer s.mux.Unlock()
		s.storage[in.Key] = in.Value
		return &pb.Result{Result: "success"}, nil
	} else {
		logger.Infof("forwarding request to %X", node.Id)
		return node.Put(ctx, in)
	}
}

func (s *ChordServer) Del(ctx context.Context, in *pb.Key) (*pb.Result, error) {
	logger := s.logger.WithFields(log.Fields{"op": "del", "key": in.Key})
	logger.Infof("request")
	hash := generate_hash(in.Key)
	node, err := s.ClosestPrecedingNode(ctx, hash)
	if err != nil {
		return nil, err
	}
	if node == nil {
		logger.Infof("key belongs to %X", s.self.Id)
		s.mux.Lock()
		defer s.mux.Unlock()
		_, ok := s.storage[in.Key]
		if !ok {
			return nil, errors.New("key not found")
		} else {
			delete(s.storage, in.Key)
			return &pb.Result{Result: "success"}, nil
		}
	} else {
		logger.Infof("forwarding request to %X", node.Id)
		return node.Del(ctx, in)
	}
}

func (s *ChordServer) Control(ctx context.Context, in *pb.ControlRequest) (*pb.Result, error) {
	if in.Control == "quit" {
	}
	if in.Control == "join" {
	}
	return &pb.Result{Result: "success"}, nil
}

func (n *ChordNode) Get(ctx context.Context, in *pb.Key) (*pb.Pair, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewDHTClient(conn)
	result, err := c.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n *ChordNode) Put(ctx context.Context, in *pb.Pair) (*pb.Result, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewDHTClient(conn)
	result, err := c.Put(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n *ChordNode) Del(ctx context.Context, in *pb.Key) (*pb.Result, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewDHTClient(conn)
	result, err := c.Del(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Serve(node *ChordNode, bootstrap_node *ChordNode, group *sync.WaitGroup) {
	logger := log.WithFields(log.Fields{"from": "serve", "id": fmt.Sprintf("%X", node.Id)})
	defer group.Done()
	server := NewChordServer(node.Address)
	lis, err := net.Listen("tcp", node.Address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChordServer(s, server)
	pb.RegisterDHTServer(s, server)
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan bool)
	go server.Serve(ctx)
	go func() {
		if err := s.Serve(lis); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
		ch <- true
	}()
	if bootstrap_node != nil {
		for {
			err := server.Join(ctx, bootstrap_node)
			if err != nil {
				logger.Warningf("%X failed to join %X at %v, retrying...", node.Id, bootstrap_node.Id, bootstrap_node.Address)
			} else {
				logger.Infof("%X joined successfully", node.Id)
				break
			}
			time.Sleep(time.Second)
		}
	}
	<-ch
	cancel()
}
