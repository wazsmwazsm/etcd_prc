package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.16.129.153:2379", "172.16.129.154:2379", "172.16.129.156:2379"},
		DialTimeout: 5 * time.Second,
		Username:    "servicetree",
		Password:    "servicetree",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	wc := cli.Watch(context.Background(), "/app/servicetree/test", clientv3.WithPrefix(), clientv3.WithPrevKV())
	for v := range wc { // 阻塞到此
		if v.Err() != nil {
			log.Fatal(err)
		}

		for _, e := range v.Events {
			log.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)
		}
	}
}
