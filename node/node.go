package node

import (
	"context"
	pb "github.com/skyzh/go-dht/protos"
	"log"
)

type DHTServer struct {
}

func (s *DHTServer) FindSuccessor(ctx context.Context, in *pb.SuccessorRequest) (*pb.SuccessorReply, error) {
	log.Printf("Received: %v", in.K)
	return &pb.SuccessorReply{Id: "23333"}, nil
}

func (s *DHTServer) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	return &pb.QueryReply{Value: "2333"}, nil
}
