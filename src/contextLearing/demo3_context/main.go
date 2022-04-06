package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 使用channel的方式实现

// 单向通道 <- 只能从里面取值
func worker(ch <-chan struct{}) {
	defer wg.Done()
LABEL:
	for {
		select {
		case <-ch:
			break LABEL
		default:
			fmt.Println("worker")
			time.Sleep(time.Second)
		}
		// 如何接收外部命令实现退出
	}

}

// make 和 new 的区别
// 同：都是用来初始化内存
// new 多用来为基本数据类型(bool\string\int\struct...)初始化内存，返回的是指针
// make 用来初始化(slice\map\channel), 返回的是对应类型

func main() {
	var exitChan = make(chan struct{})

	wg.Add(1)
	go worker(exitChan)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 5)
	exitChan <- struct{}{}
	wg.Wait() // 等待...计数器清零往下走
	fmt.Println("over")
}
