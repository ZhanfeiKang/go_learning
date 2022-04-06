package main

// kafka client demo

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	// 1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // ACK 发送数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分区 随机选出一个分区
	config.Producer.Return.Successes = true                   // 确认 成功交付的消息将在 success channel 返回

	// 2.连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config) // []string切片说明可以连多个kafka
	if err != nil {
		fmt.Println("producer closed,err: ", err)
		return
	}
	defer client.Close()

	// 3.封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("2022.4.1: this is a test log")

	// 4.发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err: ", err)
		return
	}
	fmt.Printf("pid:%v,  offset:%v\n", pid, offset)
}
