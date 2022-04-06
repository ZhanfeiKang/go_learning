package model

import (
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/shirou/gopsutil/disk"
)

type Util struct {
	Client influxdb.Client
}

type CpuInfo struct {
	CpuPercent float64 `json:"cpu_percent"`
}

type MemInfo struct {
	// Total amount of RAM on this system
	Total uint64 `json:"total"`

	// RAM available for  programs to allocate
	//
	// This value is computed from the kernel specific values.
	Available uint64 `json:"available"`

	// RAM used by programs
	//
	// This value is computed from the kernel specific values.
	Used uint64 `json:"used"`

	// Percentage of RAM used by programs
	//
	// This value is computed from the kernel specific values.
	UsedPercent float64 `json:"usedPercent"`

	Buffers uint64 `json:"buffers"`
	Cached  uint64 `json:"cached"`
}

// type UsageStat struct {
// 	Path              string  `json:"path"`
// 	Fstype            string  `json:"fstype"`
// 	Total             uint64  `json:"total"`
// 	Free              uint64  `json:"free"`
// 	Used              uint64  `json:"used"`
// 	UsedPercent       float64 `json:"used_percent"`
// 	InodesTotal       uint64  `json:"inodes_total"`
// 	InodesUsed        uint64  `json:"inodes_used"`
// 	InodesFree        uint64  `json:"inodes_free"`
// 	InodesUsedPercent float64 `json:"inodes_used_percent"`
// }

type DiskInfo struct {
	PartitionUsageStat map[string]*disk.UsageStat
}

type IOStat struct {
	BytesSent       uint64
	BytesRecv       uint64
	PacketsSent     uint64
	PacketsRecv     uint64
	BytesSentRate   float64 `json:"bytes_sent_rate"`
	BytesRecvRate   float64 `json:"bytes_recv_rate"`
	PacketsSentRate float64 `json:"packets_sent_rate"`
	PacketsRecvRate float64 `json:"packets_recv_rate"`
}

type NetInfo struct {
	// NetIOCountersStat map[string]*net.IOCountersStat
	NetIOCountersStat map[string]*IOStat
}
