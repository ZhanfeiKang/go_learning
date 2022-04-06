package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 使用context实现

func worker(ctx context.Context) {
	defer wg.Done()

	go worker2(ctx)

LABEL:
	for {
		fmt.Println("worker1")
		time.Sleep(time.Second)
		// 如何接收外部命令实现退出
		select {
		case <-ctx.Done():
			break LABEL
		default:
		}
	}
}

func worker2(ctx context.Context) {

LABEL:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		// 如何接收外部命令实现退出
		select {
		case <-ctx.Done():
			break LABEL
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 5)
	cancel()  // 调用cancel函数告诉子goroutine退出
	wg.Wait() // 等待...计数器清零往下走
	fmt.Println("cancel")
	time.Sleep(time.Second * 3)
	fmt.Println("over")
}
