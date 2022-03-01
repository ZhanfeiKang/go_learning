package main

import (
	"fmt"
	"time"
)

// write data
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println("writeData: ", i)
		time.Sleep(time.Microsecond)
	}
	close(intChan) // 关闭
}

// read data
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData: ", v)
		time.Sleep(time.Microsecond)
	}
	// readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	// time.Sleep(time.Second * 5)
	for {
		_, ok := <-exitChan // 是否能取出来
		if !ok {            // 如果读不到东西，就退出
			break
		}
	}
}
