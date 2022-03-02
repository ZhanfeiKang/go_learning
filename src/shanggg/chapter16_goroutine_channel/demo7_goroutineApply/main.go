package main

import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}

	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	// time.Sleep(time.Microsecond) //延迟1ms，等intChan有数？
	// var num int
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { // intChan 娶不到
			break
		}
		flag = true
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该 num 不是素数
				flag = false
				break
			}
		}

		if flag {
			// 将 num 放入 primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个 primeNum 协程因为取不到数据, 退出....")

	exitChan <- true
}

func main() {

	intChan := make(chan int, 8000)
	primeChan := make(chan int, 2000) // 放入结果
	exitChan := make(chan bool, 4)    // 标识退出的管道

	// start := time.Now().Unix()

	// 开启一个协程，向 intChan放入 1-8000 个数
	go putNum(intChan)

	// 开启4个协程，从 intChan 取出数据，并判断是否为素数，
	// 如果是，就放入到 primeChan

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 这里我们主线程，进行处理
	// 直接,如果没有取出4个，则主线程会一直阻塞在这里
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		// end := time.Now().Unix()
		// fmt.Println("使用协程耗时:", end-start)

		// 当我们从 exitChan 取出了4个结果，就可以放心的关闭 primeChan
		close(primeChan)
	}()

	// 遍历我们的 primeChan，把结果取出
	for {
		res, ok := <-primeChan
		// _, ok := <-primeChan
		if !ok {
			break
		}
		// 将结果输出
		fmt.Printf("素数=%d\n", res)
	}

	// 主程序结束
	fmt.Println("主程序退出...")

}
