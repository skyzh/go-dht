package main

import (
	"context"
	"github.com/skyzh/go-dht/node"
	"sync"
)

const (
	address = "127.0.0.1:50051"
)

func main() {
	group := &sync.WaitGroup{}
	group.Add(1)
	node.ServeChord(context.Background(), node.NewChordNode(address), nil, group, nil)
	group.Wait()
}
