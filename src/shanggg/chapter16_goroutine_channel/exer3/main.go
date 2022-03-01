package main

import "fmt"

func putNum(intChan chan int) {
	for i := 1; i <= 20; i++ {
		intChan <- i
	}
	close(intChan)
}

// 计算 n! ，让这个结果放入到myMap中
func calNum(intChan chan int, resChan chan map[int]int, exitChan chan bool) {

	for {
		res := 1
		m := make(map[int]int, 1)
		n, ok := <-intChan
		if !ok {
			break
		}
		for i := 1; i <= n; i++ {
			res *= i
		}
		m[n] = res
		resChan <- m
	}

	fmt.Println("有一个 calNum 协程取不到数了，结束...")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 20)
	resChan := make(chan map[int]int, 20)
	exitChan := make(chan bool, 10)

	go putNum(intChan)

	// 这里我们开启多个协程完成这个任务
	for i := 1; i <= 10; i++ {
		go calNum(intChan, resChan, exitChan)
	}

	go func() {
		for i := 0; i < 10; i++ {
			<-exitChan
		}
		close(resChan)
	}()

	for {
		m, ok := <-resChan
		if !ok {
			close(exitChan)
			break
		}

		fmt.Println(m)
	}

	fmt.Println("主线程退出...")

}
