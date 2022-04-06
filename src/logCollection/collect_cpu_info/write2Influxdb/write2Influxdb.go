package write2Influxdb

import (
	"collect_cpu_info/dao"
	"collect_cpu_info/model"
	"context"
	"log"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
)

var (
	bucket = "CPUmonitor"
	org    = "example-org"
)

// 将CPU信息写入到influxDB中
func WrtieCpuPoint(data *model.CpuInfo) (err error) {
	// 使用该方法创建一个写客户端WriteAPIBlocking并传入org和bucket参数。
	writeAPI := dao.Client.WriteAPIBlocking(org, bucket)

	p := influxdb.NewPointWithMeasurement("cpu_percent").
		AddTag("cpu", "cpu0").
		AddField("cpu_percent", data.CpuPercent).
		SetTime(time.Now())
	// 实时写入point
	err = writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Printf("add point to influxdb failed,err: %v\n", err)
		return
	}
	log.Println("insert cpu info success~")
	return
}

// 将内存信息写入到influxDB中
func WrtieMemPoint(data *model.MemInfo) (err error) {
	// 使用该方法创建一个写客户端WriteAPIBlocking并传入org和bucket参数。
	writeAPI := dao.Client.WriteAPIBlocking(org, bucket)

	p := influxdb.NewPointWithMeasurement("memory").
		AddTag("mem", "mem").
		AddField("total", data.Total).
		AddField("available", data.Available).
		AddField("used", data.Used).
		AddField("used_percent", data.UsedPercent).
		SetTime(time.Now())
	// 实时写入point
	err = writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Printf("add point to influxdb failed,err: %v\n", err)
		return
	}
	log.Println("insert mem info success~")
	return
}

// 将磁盘信息写入到influxDB中
func WrtieDiskPoint(data *model.DiskInfo) (err error) {
	// 使用该方法创建一个写客户端WriteAPIBlocking并传入org和bucket参数。
	writeAPI := dao.Client.WriteAPIBlocking(org, bucket)

	// 根据传入数据的类型插入数据
	for k, v := range data.PartitionUsageStat {
		tags := map[string]string{"path": k}
		fields := map[string]interface{}{
			"total":               v.Total,
			"free":                v.Free,
			"used":                v.Used,
			"used_percent":        v.UsedPercent,
			"inodes_total":        v.InodesTotal,
			"inodes_used":         v.InodesUsed,
			"inodes_free":         v.InodesFree,
			"inodes_used_percent": v.InodesUsedPercent,
		}
		p := influxdb.NewPoint("disk", tags, fields, time.Now())
		// 实时写入point
		err = writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Printf("add point to influxdb failed,err: %v\n", err)
			return
		}
	}

	// p := influxdb.NewPointWithMeasurement("disk").
	// 	AddTag("mem", "mem").
	// 	AddField("total", data.Total).
	// 	AddField("available", data.Available).
	// 	AddField("used", data.Used).
	// 	AddField("used_percent", data.UsedPercent).
	// 	SetTime(time.Now())

	log.Println("insert disk info success~")
	return
}

// // 将网卡信息写入到influxDB中
func WrtieNetPoint(data *model.NetInfo) (err error) {
	// 使用该方法创建一个写客户端WriteAPIBlocking并传入org和bucket参数。
	writeAPI := dao.Client.WriteAPIBlocking(org, bucket)

	// 根据传入数据的类型插入数据
	for k, v := range data.NetIOCountersStat {
		tags := map[string]string{"name": k} // 把每个网卡存为tag
		fields := map[string]interface{}{
			"bytes_sent_rate":   v.BytesSentRate,
			"bytes_recv_rate":   v.BytesRecvRate,
			"packets_sent_rate": v.PacketsSentRate,
			"packets_recv_rate": v.PacketsRecvRate,
		}
		p := influxdb.NewPoint("net", tags, fields, time.Now())
		// 实时写入point
		err = writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Printf("add point to influxdb failed,err: %v\n", err)
			return
		}
	}

	log.Println("insert net info success~")
	return
}
