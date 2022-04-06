package main

import (
	"collect_cpu_info/dao"
	"collect_cpu_info/model"
	"collect_cpu_info/write2Influxdb"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

var (
	lastNetIOStatTimeStamp int64          // 上一次获取网络IO数据的时间点
	lastNetInfo            *model.NetInfo // 上一次的网络IO数据
)

func getCpuInfo() {
	var cpuInfo = new(model.CpuInfo)
	// var cpuInfo = &model.CpuInfo{}

	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false) // 每秒统计一下cpu的使用率
	// fmt.Printf("cpu percent: %v\n", percent)

	// 写入到 influxdb中
	cpuInfo.CpuPercent = percent[0]
	write2Influxdb.WrtieCpuPoint(cpuInfo)

}

// 内存信息
func getMemInfo() {
	var memInfo = new(model.MemInfo)
	// var memInfo = &model.MemInfo{}
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed,err: %v", err)
		return
	}
	memInfo.Total = info.Total
	memInfo.Available = info.Available
	memInfo.Used = info.Used
	memInfo.UsedPercent = info.UsedPercent
	memInfo.Buffers = info.Buffers
	memInfo.Cached = info.Cached

	write2Influxdb.WrtieMemPoint(memInfo)
}

func getDiskInfo() {
	var diskInfo = &model.DiskInfo{
		PartitionUsageStat: make(map[string]*disk.UsageStat, 16),
	}
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("get disk partitions failed,err: ", err)
		return
	}

	for _, part := range parts {
		// 拿到每一个分区
		usageState, err := disk.Usage(part.Mountpoint) // 传挂载点进去
		if err != nil {
			fmt.Printf("get %s usage stat failed,err: %v", part, err)
			continue
		}
		diskInfo.PartitionUsageStat[part.Mountpoint] = usageState
	}

	write2Influxdb.WrtieDiskPoint(diskInfo)
}

// 网卡信息
func getNetInfo() {
	var netInfo = &model.NetInfo{
		NetIOCountersStat: make(map[string]*model.IOStat, 8),
	}
	currentTimeStamp := time.Now().Unix()
	netIOs, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net io counters failed, err:%v", err)
		return
	}
	for _, netIO := range netIOs {
		var ioStat = new(model.IOStat)
		ioStat.BytesSent = netIO.BytesSent
		ioStat.BytesRecv = netIO.BytesRecv
		ioStat.PacketsSent = netIO.PacketsSent
		ioStat.PacketsRecv = netIO.PacketsRecv
		// 将具体网卡数据的ioStat变量添加到map中
		netInfo.NetIOCountersStat[netIO.Name] = ioStat // 不要放到continue下面
		// 开始计算网卡相关速率
		if lastNetIOStatTimeStamp == 0 || lastNetInfo == nil {
			continue
		}
		// 计算时间间隔
		interval := currentTimeStamp - lastNetIOStatTimeStamp
		// 计算速率
		ioStat.BytesSentRate = (float64(ioStat.BytesSent) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].BytesSent)) / float64(interval)
		ioStat.BytesRecvRate = (float64(ioStat.BytesRecv) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].BytesRecv)) / float64(interval)
		ioStat.PacketsSentRate = (float64(ioStat.PacketsSent) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].PacketsSent)) / float64(interval)
		ioStat.PacketsRecvRate = (float64(ioStat.PacketsRecv) - float64(lastNetInfo.NetIOCountersStat[netIO.Name].PacketsRecv)) / float64(interval)

	}
	// 更新全局记录的上一次采集的网卡时间点和网卡数据
	lastNetIOStatTimeStamp = currentTimeStamp // 更新时间
	lastNetInfo = netInfo
	// 发送到influxdb
	write2Influxdb.WrtieNetPoint(netInfo)
}

func run(interval time.Duration) {

	// ticker := time.Tick(time.Second) //返回一个通道，每一秒给一个信号
	ticker := time.Tick(interval)
	for range ticker {
		getCpuInfo()
		getMemInfo()
		getDiskInfo()
		getNetInfo()
		fmt.Println("---")
	}
}

func main() {
	dao.InitConnInflux()
	defer dao.CloseConn()

	run(time.Second * 2) // 每一秒钟收集一次
}
