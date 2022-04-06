package main

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influxdb-client-go/api"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
)

// influx demo
/*
kkite  kkitekkite | org = "example-org" | bucket = "example-bucket" | token  = "example-token"
*/
var (
	bucket = "mybucket"
	org    = "example-org"
	token  = "HtP0oGo-MlXkYT5071pGkEf7cfWAf2CR_l-pOKIJoYkQwInD-1jliGTNEbnaoPgS0edR-rHeG-LRVLRTdpqAOg=="
	// Store the URL of your InfluxDB instance
	url = "http://127.0.0.1:8086"
)

// connect
func connInflux() influxdb.Client {

	// 创建 InfluxDB Go 客户端并传入url和token参数
	// client := influxdb.NewClient(url, token, influxdb.DefaultOptions().SetBatchSize(20))
	client := influxdb.NewClient(url, token)
	return client
}

/*
WriteAPIBlocking: 阻塞
WriteAPI: 非阻塞。即异步写
			异步写的机制：数据首先被异步写入到一个buffer，满足一定条件时才会写入数据库。条件为：
						- 要么缓存数据达到5000条（batch size）
						- 要么等待定时flush，默认一秒一次
*/
// write
func wrtiePoint(client influxdb.Client) (err error) {
	// 使用该方法创建一个写客户端WriteAPIBlocking并传入org和bucket参数。
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// 创建一个 point 并使用API writer struct 的方法将其写入 influxdb
	// tags := map[string]string{"cpu": "ih-cpu"}
	// fields := map[string]interface{}{}

	// p := influxdb.NewPoint("cpu_usage", tags, fields, time.Now())
	p := influxdb.NewPointWithMeasurement("state").
		AddTag("unit", "temperature").
		AddField("avg", 26.6).
		AddField("max", 42.0).
		SetTime(time.Now())
	// 实时写入point
	err = writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("add point to influxdb failed,err: %v\n", err)
		return
	}
	fmt.Println("insert ")
	return
}

// writeCPU

// query
func queryDB(client influxdb.Client) (res *api.QueryTableResult, err error) {
	// 要查询数据，请创建一个 InfluxDB查询客户端并传入您的 InfluxDB org
	queryAPI := client.QueryAPI(org)
	// 第1行：指定数据源
	// 第2行：限定时间区间，flux查询必须限定时间区间
	// 		  如果 stop 省略，就默认是当前时间
	// |>  ：类似于linux管道，即将上一步的查询结果，传递给下一步操作
	// filter：对每行 record 进行过滤
	// fn ：function 的缩写，类似 js 中的匿名函数，这里非常像箭头函数
	// r ：record 的缩写
	//
	result, err := queryAPI.Query(context.Background(), `from(bucket:"example-bucket")
    |> range(start: -1h) 
    |> filter(fn: (r) => r._measurement == "state")`)
	if err != nil {
		fmt.Println("query from influxdb failed,err:", err)
		return nil, err
	}

	for result.Next() {
		// TableChanged：标识组密钥何时更改。
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// Record：返回最后解析的 FluxRecord 并允许访问 value 和 row 属性。
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	if result.Err() != nil {
		fmt.Printf("query parsing error:%s\n", result.Err().Error())
	}
	return
}

func main() {
	conn := connInflux()
	defer conn.Close()

	// insert
	wrtiePoint(conn)
	// queryDB(conn)

	// 1.7
	// fmt.Sprintf("SELECT * FROM %s LIMIT %d","cpu_usage",10)
	// tags：	"cpu"   : "ih-cpu"
	// fields:	"idle"  : 201.1
	// 			"system": 43.3
	// 			"user"  : 86.6
}
