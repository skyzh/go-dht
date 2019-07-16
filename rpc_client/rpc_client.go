package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

const (
	N = 50
)

func main() {
	for key := 1000; key <= 2000; key += 100 {
		target := rand.Intn(N)
		address := fmt.Sprintf("127.0.0.1:%d", 40000+target)
		log.Infof("put key %v at %v", key, address)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewDHTClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err = c.Put(ctx, &pb.Pair{Key: fmt.Sprintf("%d", key), Value: fmt.Sprintf("%d", 23330000 + key)})
		if err != nil {
			log.Warnf("could not put: %v", err)
		}
		for i := 0; i < 10; i++ {
			target := rand.Intn(N)
			address := fmt.Sprintf("127.0.0.1:%d", 40000+target)
			log.Infof("request key %v at %v", key, address)
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewDHTClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			p, err := c.Get(ctx, &pb.Key{Key: fmt.Sprintf("%d", key)})
			if err != nil {
				log.Warnf("could not get: %v", err)
			} else {
				log.Infof("value: %v", p.Value)
			}
		}
	}

}
