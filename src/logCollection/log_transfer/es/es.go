package es

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// 将日志数据写入ElasticSearch

type ESClient struct {
	client      *elastic.Client
	index       string
	logDataChan chan interface{}
}

var (
	esClient *ESClient
)

func Init(addr, index string, goroutineNum, maxSize int) (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + addr))
	if err != nil {
		// Handle error
		panic(err)
	}
	esClient = &ESClient{
		client:      client,
		index:       index,
		logDataChan: make(chan interface{}, maxSize),
	}

	fmt.Println("connect to es success~")

	// 从通道取出数据，写入到kafka中去
	for i := 0; i < goroutineNum; i++ {
		go sendToES()
	}

	return
}

func sendToES() {
	for m1 := range esClient.logDataChan {
		// b, err := json.Marshal(m1)
		// if err != nil {
		// 	fmt.Println("marshal m1 failed,err: ", err)
		// 	continue
		// }
		put, err := esClient.client.Index().
			Index(esClient.index).
			// Id("3").
			BodyJson(m1).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Indexed user %s to index %s, type %s\n", put.Id, put.Index, put.Type)
	}
}

// 通过一个首字母大写的函数从包外接收msg，发送到chan中
func PutLogData(msg interface{}) {
	esClient.logDataChan <- msg
}
