package main

import (
	"context"
	pb "github.com/skyzh/go-dht/chord"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	address = "localhost:50051"
	port    = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) FindSuccessor(ctx context.Context, in *pb.SuccessorRequest) (*pb.SuccessorReply, error) {
	log.Printf("Received: %v", in.K)
	return &pb.SuccessorReply{Id: "23333"}, nil
}

func make_request() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FindSuccessor(ctx, &pb.SuccessorRequest{K: "test"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Id)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	go make_request()
	pb.RegisterChordServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
