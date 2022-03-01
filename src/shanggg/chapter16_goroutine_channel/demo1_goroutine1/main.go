package main

import (
	"fmt"
	"strconv"
	"time"
)

// 在主线程（可以理解为进程）中，开启一个goroutine，该协程每隔1秒输出"hello kkite~"
// 在主线程中也每隔1秒输出"hello,daisy~~"，输出10次后退出程序
// 要求主线程和goroutine同时执行

// 编写一个函数，没隔1秒输出"hello kkite~"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test(): hello kkite~" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test() // 开启了一个协程

	for i := 1; i <= 10; i++ {
		fmt.Println("main(): hello daisy~~" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
