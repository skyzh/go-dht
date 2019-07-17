#/bin/bash

go run ./rpc_cluster &
sleep 10
go run ./rpc_client
pkill rpc_cluster
