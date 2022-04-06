package kafka

// kafka相关操作

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	client  sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage
)

// Init 是初始化全局的kafka Client
func Init(address []string, chanSize int64) (err error) {
	// 1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // ACK 发送数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分区 随机选出一个分区
	config.Producer.Return.Successes = true                   // 确认 成功交付的消息将在 success channel 返回

	// 2.连接kafka
	client, err = sarama.NewSyncProducer(address, config) // []string切片说明可以连多个kafka
	if err != nil {
		logrus.Error("kafka:producer closed,err: ", err)
		return
	}

	// 初始化MsgChan
	msgChan = make(chan *sarama.ProducerMessage, chanSize)
	// 起一个后台的goroutine从msgChan中读数据
	go sendMsg()
	return
}

// 从 MsgChan 中读取 msg ，发送给kafka
func sendMsg() {
	for {
		select {
		case msg := <-msgChan:
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				logrus.Warning("send msg failed,err: ", err)
				return
			}
			logrus.Infof("send msg to kafka success~ pid:%v offset:%v", pid, offset)
		}
	}
}

// 定义一个函数，向外暴露msgChan,只能向里面发送值
func ToMsgChan(msg *sarama.ProducerMessage) {
	msgChan <- msg
}
