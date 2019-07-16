package node

import (
	"context"
	log "github.com/sirupsen/logrus"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

func (s *ChordServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.Result, error) {
	return &pb.Result{Result: "success"}, nil
}

func (s *ChordServer) Put(ctx context.Context, in *pb.PutRequest) (*pb.Result, error) {
	return nil, nil
}

func (s *ChordServer) Del(ctx context.Context, in *pb.DelRequest) (*pb.Result, error) {
	return nil, nil
}

func Serve(node *ChordNode, bootstrap_node *ChordNode, group *sync.WaitGroup) {
	defer group.Done()
	server := NewChordServer(node.Address)
	lis, err := net.Listen("tcp", node.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChordServer(s, server)
	pb.RegisterDHTServer(s, server)
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan bool)
	go server.Serve(ctx)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		ch <- true
	}()
	if bootstrap_node != nil {
		for {
			err := server.Join(ctx, bootstrap_node)
			if err != nil {
				log.Warningf("%X failed to join %X at %v, retrying...", node.Id, bootstrap_node.Id, bootstrap_node.Address)
			} else {
				log.Infof("%X joined successfully", node.Id)
				break
			}
			time.Sleep(time.Second)
		}
	}
	<-ch
	cancel()
}
