package dao

import (
	influxdb "github.com/influxdata/influxdb-client-go/v2"
)

var (
	Client influxdb.Client

	token = "HtP0oGo-MlXkYT5071pGkEf7cfWAf2CR_l-pOKIJoYkQwInD-1jliGTNEbnaoPgS0edR-rHeG-LRVLRTdpqAOg=="
	// Store the URL of your InfluxDB instance
	url = "http://127.0.0.1:8086"
)

// connect
func InitConnInflux() {

	// 创建 InfluxDB Go 客户端并传入url和token参数
	// client := influxdb.NewClient(url, token, influxdb.DefaultOptions().SetBatchSize(20))
	Client = influxdb.NewClient(url, token)
}

func CloseConn() {
	Client.Close()
}
