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
	log.SetLevel(log.InfoLevel)
	logger := log.WithFields(log.Fields{"from": "main"})
	group := &sync.WaitGroup{}
	join_group := &sync.WaitGroup{}
	var nodes []*node.ChordNode
	for i := 0; i < N; i++ {
		addr := fmt.Sprintf("127.0.0.1:%d", 40000+i)
		nodes = append(nodes, node.NewChordNode(addr))
	}
	for i := 0; i < N; i++ {
		group.Add(1)
		join_group.Add(1)
		if i == 0 {
			logger.Tracef("initial node %X listening at %v", nodes[i].Id, nodes[i].Address)
			logger.Infof("initial node has been set up")
			go node.ServeChord(nodes[i], nil, group, join_group)
		} else {
			logger.Tracef("%X listening at %v bootstrapped with %X", nodes[i].Id, nodes[i].Address, nodes[0].Id)
			go node.ServeChord(nodes[i], nodes[0], group, join_group)
		}
	}
	logger.Infof("all nodes have been set up")
	go func() {
		time.Sleep(time.Second)
	} ()
	join_group.Wait()
	logger.Infof("all nodes have joined network")
	group.Wait()
}
