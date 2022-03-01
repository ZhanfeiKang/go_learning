package main

import "fmt"

func main() {

	intChan := make(chan int, 3)

	intChan <- 100
	intChan <- 200
	close(intChan) // close
	// 这是不能够再写入数据到channel
	// intChan <- 300
	fmt.Println("ook~")

	// 当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1: ", n1)

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 // 放入100个数据到管道
	}

	// 遍历管道不能使用普通的for循环
	// 遍历
	// for i := 0; i < len(intChan2); i++ {}	//错误遍历

	// 在遍历时，如果channel没有关闭，则会出现deadlock错误

	// 在遍历时，若谷channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v: ", v)
	}
}
