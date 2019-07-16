package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	pb "github.com/skyzh/go-dht/protos"
	"google.golang.org/grpc"
	"math/rand"
	"sync"
	"time"
)

const (
	N               = 50
	CONCURRENT_JOBS = 10
	TOTAL_KEYS      = 2000
)

func generate_pairs() map[string]string {
	data := make(map[string]string)
	for i := 0; i < TOTAL_KEYS; i++ {
		k := fmt.Sprintf("%016X", rand.Uint64())
		v := fmt.Sprintf("%016X", rand.Uint64())
		data[k] = v
	}
	return data
}

var cnt = 0

func put_pair(addr, k, v string, group *sync.WaitGroup, concurrent_jobs chan bool) {
	defer group.Done()
	concurrent_jobs <- true
	defer func() { <-concurrent_jobs }()
	cnt = cnt + 1
	log.Infof("(%4d/%4d) putting %v into %v...", cnt, TOTAL_KEYS, k, addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection to %v failed: %v", addr, err)
	}
	defer conn.Close()
	c := pb.NewDHTClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	_, err = c.Put(ctx, &pb.Pair{Key: k, Value: v})
	if err != nil {
		log.Warnf("could not put to %v: %v", addr, err)
	}
}

func query_key(addr, k, v string, group *sync.WaitGroup, concurrent_jobs chan bool, should_error bool) {
	defer group.Done()
	concurrent_jobs <- true
	defer func() { <-concurrent_jobs }()
	cnt = cnt + 1
	log.Infof("(%4d/%4d) querying %v from %v...", cnt, TOTAL_KEYS, k, addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection to %v failed: %v", addr, err)
	}
	defer conn.Close()
	c := pb.NewDHTClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	val, err := c.Get(ctx, &pb.Key{Key: k})
	if err != nil {
		if !should_error {
			log.Warnf("could not query from %v: %v", addr, err)
		}
	} else {
		if should_error {
			log.Errorf("expected %v to be deleted", k)
		} else if val.Value != v {
			log.Errorf("(got) %v != (expected) %v", val.Value, v)
		}
	}
}

func del_key(addr, k, v string, group *sync.WaitGroup, concurrent_jobs chan bool) {
	defer group.Done()
	concurrent_jobs <- true
	defer func() { <-concurrent_jobs }()
	cnt = cnt + 1
	log.Infof("(%4d/%4d) delete %v from %v...", cnt, TOTAL_KEYS, k, addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection to %v failed: %v", addr, err)
	}
	defer conn.Close()
	c := pb.NewDHTClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	_, err = c.Del(ctx, &pb.Key{Key: k})
	if err != nil {
		log.Warnf("could not delete from %v: %v", addr, err)
	}
}

func main() {
	log.Info("generate pairs...")
	pairs := generate_pairs()
	log.Infof("%d pairs generated", len(pairs))
	log.Info("putting pairs into DHT...")
	group := &sync.WaitGroup{}
	concurrent_jobs := make(chan bool, CONCURRENT_JOBS)
	cnt = 0
	for k, v := range pairs {
		time.Sleep(time.Millisecond * 5)
		target := rand.Intn(N)
		address := fmt.Sprintf("127.0.0.1:%d", 40000+target)
		go put_pair(address, k, v, group, concurrent_jobs)
		group.Add(1)
	}
	group.Wait()
	log.Info("waiting for the network to be stable...")
	time.Sleep(time.Second * 5)

	log.Info("querying pairs from DHT...")
	cnt = 0
	for k, v := range pairs {
		time.Sleep(time.Millisecond * 5)
		target := rand.Intn(N)
		address := fmt.Sprintf("127.0.0.1:%d", 40000+target)
		go query_key(address, k, v, group, concurrent_jobs, false)
		group.Add(1)
	}
	group.Wait()
	log.Info("waiting for the network to be stable...")
	time.Sleep(time.Second * 5)

	log.Info("deleting pairs from DHT...")
	cnt = 0
	for k, v := range pairs {
		time.Sleep(time.Millisecond * 5)
		target := rand.Intn(N)
		address := fmt.Sprintf("127.0.0.1:%d", 40000+target)
		go del_key(address, k, v, group, concurrent_jobs)
		group.Add(1)
	}
	group.Wait()
	log.Info("waiting for the network to be stable...")
	time.Sleep(time.Second * 5)

	log.Info("querying pairs from DHT...")
	cnt = 0
	for k, v := range pairs {
		time.Sleep(time.Millisecond * 5)
		target := rand.Intn(N)
		address := fmt.Sprintf("127.0.0.1:%d", 40000+target)
		go query_key(address, k, v, group, concurrent_jobs, true)
		group.Add(1)
	}
	group.Wait()

	log.Info("testing sequence ends, all tests passed.")
}
