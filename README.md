# go-dht

[![Build Status](https://travis-ci.com/skyzh/go-dht.svg?branch=master)](https://travis-ci.com/skyzh/go-dht)

DHT models in golang. Use gRPC for communication.


## Protocol

gRPC is used for communication between nodes. `Chord` and `Kad` RPCs are defined as
two different services in protobuf file. And here's an additional service called `DHT`,
through which clients control DHT node and modify keys in DHT.

## Chord

Chord is a DHT protocol and algorithm. Here I use SHA-1 as hashing function, and select first 
16 bits of hashing result as node identifier. You may change it in `node/chord.go`.

To run tests,
```bash
go test ./node -v -short
```

For system test (build cluster and stabilize it, usually takes longer time),
```bash
go test ./node -v
```

To quickly setup a cluster of 50 nodes,
```bash
go run ./rpc_cluster
```

To run a single node,
```bash
go run ./rpc_node
```

To test key modification in cluster (put, query and delete 2000 entries),
```bash
go run ./rpc_client
```

## Kademlia

Kademlia is expected to be implemented.
