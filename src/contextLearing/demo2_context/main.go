package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 全局变量的方式
var exit bool

func worker() {
	defer wg.Done()
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		// 如何接收外部命令实现退出
		if exit {
			break
		}
	}

}

func main() {
	wg.Add(1)
	go worker()
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 5)
	exit = true // 通过修改全局变量的方式通知子goroutine退出
	wg.Wait()   // 等待...计数器清零往下走
	fmt.Println("over")
}
