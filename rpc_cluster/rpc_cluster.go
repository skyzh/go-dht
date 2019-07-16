package main

import (
	"fmt"
	"github.com/skyzh/go-dht/node"
	"log"
	"sync"
)

func main() {
	group := &sync.WaitGroup{}
	N := 50
	var nodes []*node.ChordNode
	for i := 0; i < N; i++ {
		addr := fmt.Sprintf("127.0.0.1:%d", 40000+i)
		nodes = append(nodes, node.NewChordNode(addr))
	}
	for i := 0; i < N; i++ {
		group.Add(1)
		log.Printf("%X listening at %v", nodes[i].Id, nodes[i].Address)
		if i == 0 {
			go node.Serve(nodes[i], nil, group)
		} else {
			go node.Serve(nodes[i], nodes[0], group)
		}
	}
	group.Wait()
}
