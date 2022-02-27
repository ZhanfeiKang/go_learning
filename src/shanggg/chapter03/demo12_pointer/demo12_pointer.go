package main

import "fmt"

func main() {
	// 基本数据类型在内存布局
	var i int = 10
	// i 的地址是什么，&i
	fmt.Println("i的地址=", &i)

	// 下面的 var ptr *int = &i
	// 1. ptr是一个指针变量
	// 2. ptr的类型 *int
	// 3. ptr本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Println("ptr 的地址=", &ptr)
	fmt.Printf("ptr 指向的值=%v\n", *ptr)

	var num int = 9
	var ptr2 *int = &num

	*ptr2 = 10

	fmt.Println(num)
}
