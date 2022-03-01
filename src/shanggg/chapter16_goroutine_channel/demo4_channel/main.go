package main

import "fmt"

func main() {
	//演示一下管道的使用
	// 1.创建一个可以存放3个int类型的管道
	// var intChan chan int
	intChan := make(chan int, 3)

	// 2.看看intChan是什么
	fmt.Println("intChan 的值: ", intChan)
	fmt.Println("intChan 本身的地址: ", &intChan)

	// 3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 43
	// intChan <- 45
	// 注意：当我们给管道写入数据时，不能超过其容量

	// 4.看看管道的长度和cap(容量)
	// 容量不能自动增长
	fmt.Println("channel len: ", len(intChan))
	fmt.Println("channel cap: ", cap(intChan))

	// 5.从管道中读取数据

	// var num2 int
	num2 := <-intChan
	fmt.Println("num2: ", num2)
	fmt.Println("channel len: ", len(intChan))
	fmt.Println("channel cap: ", cap(intChan))

	// 6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock

	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3: ", num3, ", num4: ", num4)

}
