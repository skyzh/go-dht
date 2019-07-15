package node

import (
	"context"
	pb "github.com/skyzh/go-dht/protos"
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
