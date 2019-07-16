package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/skyzh/go-dht/node"
	"sync"
	"time"
)

const (
	N = 50
)

func main() {
	logger := log.WithFields(log.Fields{"from": "main"})
	group := &sync.WaitGroup{}
	var nodes []*node.ChordNode
	for i := 0; i < N; i++ {
		addr := fmt.Sprintf("127.0.0.1:%d", 40000+i)
		nodes = append(nodes, node.NewChordNode(addr))
	}
	for i := 0; i < N; i++ {
		group.Add(1)
		if i == 0 {
			logger.Infof("initial node %X listening at %v", nodes[i].Id, nodes[i].Address)
			go node.Serve(nodes[i], nil, group)
		} else {
			logger.Infof("%X listening at %v bootstrapped with %X", nodes[i].Id, nodes[i].Address, nodes[0].Id)
			go node.Serve(nodes[i], nodes[0], group)
		}
	}
	go func() {
		time.Sleep(time.Second)
	} ()
	group.Wait()
}
