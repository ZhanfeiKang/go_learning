package main

import (
	"fmt"
	"log_transfer/es"
	"log_transfer/kafka"
	"log_transfer/model"

	"github.com/go-ini/ini"
)

// log transfer
// 从 kafka 消费日志数据，写入 ES

func main() {
	// 1.加载配置文件
	var cfg = new(model.Config)
	err := ini.MapTo(cfg, "./config/logtransfer.ini")
	if err != nil {
		fmt.Println("load config failed! err: ", err)
		panic(err)
	}
	fmt.Printf("%#v\n", *cfg)
	fmt.Println("load config success...")
	// 2.连接ES
	err = es.Init(cfg.ESConf.Address, cfg.ESConf.Index, cfg.ESConf.GoNum, cfg.ESConf.MaxSize)
	if err != nil {
		fmt.Println("init es failed! err: ", err)
		panic(err)
	}
	fmt.Println("Init ES kafka success...")

	// 3.连接kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Println("Init kafka failed! err: ", err)
		panic(err)
	}
	fmt.Println("Init kafka success...")

	// 让程序在这停顿！
	// select不占cpu
	select {}
}
