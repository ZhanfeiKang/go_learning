package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// 代码连接etcd

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v", err)
		return
	}
	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	str := `[{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/s4.log","topic":"s4"},{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/web.log","topic":"web_log"}]`
	// str := `[{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/s4.log","topic":"s4"},
	// 		{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/web.log","topic":"web_log"},
	// 		{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/kkite.log","topic":"kkite_log"}]`
	// str := `[{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/s4.log","topic":"s4"},
	// 		{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/web.log","topic":"web_log"},
	// 		{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/kkite.log","topic":"kkite_log"},
	// 		{"path":"D:/ProgramFiles/kafka_2.12-3.1.0/logs/daisy.log","topic":"daisy_log"}]`
	_, err = cli.Put(ctx, "collect_log_10.136.232.30_conf", str)
	// _, err = cli.Put(context.TODO(), "daisy", "beautiful")
	if err != nil {
		fmt.Printf("put to etcd failed,err:%v", err)
		return
	}
	cancel()

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := cli.Get(ctx, "collect_log_10.136.232.30_conf")
	if err != nil {
		fmt.Printf("get from etcd failed,err:%v", err)
		return
	}
	for _, ev := range gr.Kvs {
		fmt.Printf("key:%s  value:%s\n", ev.Key, ev.Value)
	}
	cancel()
}
