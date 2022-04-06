package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// 获取cpu信息

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed,err: %v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false) // 每秒统计一下cpu的使用率
		fmt.Printf("cpu percent: %v\n", percent)
	}
}

// CPU 负载
func getLoad() {
	info, err := load.Avg()
	if err != nil {
		fmt.Printf("load.Avg() failed,err: %v", err)
		return
	}
	fmt.Println(info)
}

// 内存信息
func getMemInfo() {
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed,err: %v", err)
		return
	}
	fmt.Println(info)
}

// host info
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v\n uptime:%v\n boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// 磁盘 信息
func getDiskInfo() {
	// 获取所有分区信息
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("get disk partitions failed,err: ", err)
		return
	}
	// fmt.Println(parts)
	for _, part := range parts {
		partInfo, err := disk.Usage(part.Mountpoint) // 传一个挂载点或者盘符的信息 mountpoint
		if err != nil {
			fmt.Println("get part stat failed,err: ", err)
			return
		}
		fmt.Println(partInfo)
	}

	// 磁盘IO
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

// 网络信息
func getNetInfo() {
	ioCounters, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net io counters failed,err:%v", err)
		return
	}
	// fmt.Println(ioCounters)
	for _, netIO := range ioCounters {
		fmt.Println(netIO)
	}
}

func main() {

	// getLoad()
	// getMemInfo()
	// getHostInfo()
	// getDiskInfo()
	getNetInfo()
	// getCpuInfo()
}
