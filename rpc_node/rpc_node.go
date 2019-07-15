package main

import (
	pb "github.com/skyzh/go-dht/protos"
	"github.com/skyzh/go-dht/node"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	address = "127.0.0.1:50051"
)

func main() {
	server := &node.DHTServer{}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChordServer(s, server)
	pb.RegisterDHTServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
