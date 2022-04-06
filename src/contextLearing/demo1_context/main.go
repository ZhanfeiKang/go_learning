package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 初始的例子

func worker() {
	defer wg.Done()
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出

}

func main() {
	wg.Add(1)
	go worker()
	// 如何优雅的实现结束子goroutine
	wg.Wait() // 等待...计数器清零往下走
	fmt.Println("over")
}
