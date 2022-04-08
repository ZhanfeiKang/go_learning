package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
	err    error
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func getVersion() {
	esversion, err := client.ElasticsearchVersion("http://192.168.56.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}

func main() {
	client, err = elastic.NewClient(elastic.SetURL("http://192.168.56.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success~")

	getVersion()
	// esversion, err := client.ElasticsearchVersion("http://192.168.56.1:9200")
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// fmt.Printf("Elasticsearch version %s\n", esversion)

	// p1 := Person{Name: "kkite", Age: 22, Married: false}
	// p2 := Person{Name: "daisy", Age: 3, Married: false}
	p3 := Person{Name: "zhangsan", Age: 52, Married: true}
	put1, err := client.Index().
		Index("user").
		Id("3").
		BodyJson(p3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
