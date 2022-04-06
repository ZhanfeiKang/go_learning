package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"logagent/common"
	"logagent/tailfile"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
)

// etcd 相关操作

var (
	client *clientv3.Client
)

func Init(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v", err)
		return
	}
	return
}

// 拉取日志收集配置项的函数
func GetConf(key string) (collectEntryList []common.CollectEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		logrus.Errorf("get conf from etcd by key:%s failed,err:%v", key, err)
		return
	}
	if len(resp.Kvs) == 0 {
		logrus.Warningf("get len:0 conf from etcd by key:%s", key)
		return
	}
	ret := resp.Kvs[0]
	// ret.Value // json格式字符串
	fmt.Println(ret.Value)
	err = json.Unmarshal(ret.Value, &collectEntryList)
	if err != nil {
		logrus.Errorf("json unmarshal failed,err:%v", err)
		return
	}
	return
}

// 监控 etcd 中日志收集项配置变化的函数
func WatchConf(key string) {
	for {
		watchCh := client.Watch(context.Background(), key)

		// 通道没值会阻塞
		for wresp := range watchCh {

			logrus.Info("get new conf from etcd~")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%v key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				var newConf []common.CollectEntry

				if evt.Type == clientv3.EventTypeDelete {
					// 如果是删除
					logrus.Warning("warnning:etcd delete the key!!!")
					tailfile.SendNewConf(newConf) // 没有人接收就是阻塞的
					continue
				}
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					logrus.Errorf("json unmarshal new conf failed,err:%v", err)
					continue
				}
				// 告诉tailfile这个模块应该启用新的配置了！
				tailfile.SendNewConf(newConf) // 没有人接收就是阻塞的
			}
		}
	}
}
