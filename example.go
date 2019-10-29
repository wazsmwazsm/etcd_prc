package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.16.129.153:2379", "172.16.129.154:22379", "172.16.129.156:32379"},
		DialTimeout: 5 * time.Second,
		Username:    "servicetree",
		Password:    "servicetree",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	errHandler := func(err error) {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	_, err = cli.Put(context.TODO(), "/app/servicetree/test", "hello my friend!")
	if err != nil {
		errHandler(err)
	}

	getResp, err := cli.Get(context.TODO(), "/app/servicetree/test")
	if err != nil {
		errHandler(err)
	}
	for _, ev := range getResp.Kvs {
		log.Printf("%s : %s\n", ev.Key, ev.Value)
	}

}