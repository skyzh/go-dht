package main

import "github.com/skyzh/go-dht/node"

const (
	address = "127.0.0.1:50051"
)

func main() {
	node.Serve(address)
}
