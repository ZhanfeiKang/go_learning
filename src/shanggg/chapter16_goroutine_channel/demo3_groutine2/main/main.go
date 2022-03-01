package main

import (
	"fmt"
	"sync"
	"time"
)

// 计算1-200的各个数的阶乘，并且把各个数的结成放入到map中
// 最后显示出来。要求使用groutine完成

// 思路
// 1.编写一个函数，来计算各个数阶乘，并放入到map中。
// 2.我们启动的协程多个，统计的结果将放入到map中
// 3.map应该是全局的

var (
	myMap = make(map[int]int, 10)

	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁（写锁）
	// sync 是包：synchornized 同步
	// Mutex : 是互斥
	lock sync.Mutex
)

// test函数就是计算 n! ，让这个结果放入到myMap中
func test(n int) {

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 这里我们将res放入到myMap中	[问题1：并发问题]
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {

	// 这里我们开启多个协程完成这个任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	// 休眠10s钟
	time.Sleep(time.Second * 5) // [问题2：等多久]

	// 这里我们输出结果，遍历这个结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]:%d\n", i, v)
	}
	lock.Unlock()
}
