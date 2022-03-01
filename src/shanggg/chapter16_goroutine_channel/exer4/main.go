package main

import (
	"fmt"
)

/*
	1.启动一个协程，将1-2000的数放入到一个channel中，比如 numChan
	2.启动8个协程，从numChan取出数（比如n），并计算1+..+n的值，并存放到 resChan
	3.最后8个协程同完成工作后，再遍历 resChan, 显示结果[ 如 res[1]=1..res[10]=55..]
	4.注意：resChan chan int 不合适
*/
type NAndSum struct {
	N   int
	Sum int
}

func putNum(numChan chan int) {
	for i := 1; i <= 8000; i++ {
		numChan <- i
	}
	close(numChan)
}

func calculate(numChan chan int, exitChan chan bool, resChan chan NAndSum) {

	// time.Sleep(time.Second)
	for {
		var sum int
		num, ok := <-numChan
		if !ok { // 取不到
			break
		}

		for i := 1; i <= num; i++ {
			sum += i
		}

		ns := NAndSum{num, sum}
		resChan <- ns

	}

	fmt.Println("有一个 calculate 协程由于取不到数，结束...")
	exitChan <- true

}

func main() {
	numChan := make(chan int, 2000)
	exitChan := make(chan bool, 8)
	resChan := make(chan NAndSum, 2000)

	go putNum(numChan)

	for i := 0; i < 8; i++ {
		go calculate(numChan, exitChan, resChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resChan)
	}()

	for {
		ns, ok := <-resChan
		if !ok {
			close(exitChan)
			break
		}
		fmt.Println(ns)
	}

	fmt.Println("主线程退出...")

}
