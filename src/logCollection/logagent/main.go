package main

import (
	"fmt"
	"logagent/common"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/tailfile"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// 日志收集的客户端
// 类似的开源项目还有filebeat
// 收集指定目录下的日志文件，发送到kafka中

// 现在的技能包：
// 往kafka发数据
// 使用tail读日志文件

// 整个logagent的配置结构体
type Config struct {
	KafkaConfig `ini:"kafka"`
	Collect     `ini:"collect"`
	EtcdConfig  `ini:"etcd"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	ChanSize int64  `ini:"chan_size"`
}

type Collect struct {
	LogFilePath string `ini:"logfile_path"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

func run() {
	select {}
}

func main() {
	// -1：获取本机ip，为后续etcd取配置文件打下基础
	ip, err := common.GetOutboundIP()
	if err != nil {
		logrus.Errorf("get ip failed,err:%v", err)
		return
	}
	var configObj = new(Config)
	// 0.读配置文件 `go-ini`
	// cfg, err := ini.Load("./conf/config.ini")
	// if err != nil {
	// 	logrus.Error("load config failed,err:%v", err)
	// 	return
	// }
	// kafkaAddr := cfg.Section("kafka").Key("address").String()
	// fmt.Println(kafkaAddr)
	err = ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		logrus.Errorf("load config failed,err:%v", err)
		return
	}
	fmt.Printf("%#v\n", configObj)
	// 1.初始化连接kafka（做好准备工作）
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Errorf("init kafka failed,err:%v", err)
		return
	}
	logrus.Info("init kafka success~")

	// 初始化etcd连接
	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		logrus.Errorf("init etcd failed,err:%v", err)
		return
	}
	// 从etcd中拉取要收集日志的配置项
	collectKey := fmt.Sprintf(configObj.EtcdConfig.CollectKey, ip)
	allConf, err := etcd.GetConf(collectKey)
	if err != nil {
		logrus.Errorf("get conf from etcd failed,err:%v", err)
		return
	}
	fmt.Println(allConf)
	// 派一个小弟去监控etcd中 collectKey 对应值的变化
	go etcd.WatchConf(collectKey)
	// 2.根据配置中的日志路径初始化tail，使用tail去收集日志
	// 有一个收集项，就启动一个tailObj去收集它
	err = tailfile.Init(allConf) // 从etcd中获取的配置项传到init中
	if err != nil {
		logrus.Errorf("init tailfile failed,err:%v", err)
		return
	}
	logrus.Info("init tailfile success~")

	run()
}
