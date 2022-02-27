package main

import "fmt"

func main() {
	// 演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}

	// 声明/定义一个切片
	// slice := intArr[1:3]
	// 1. slice 就是切片名
	// 2. intArr[1:3] 表示 slice 引用到intArr这个数组
	// 3. [1,3)
	slice := intArr[1:3]
	fmt.Println("intArr :", intArr)
	fmt.Println("slice 的元素是 :", slice)
	fmt.Println("slice 的元素个数是 :", len(slice))
	fmt.Println("slice 的容量是 :", cap(slice)) // 切片的容量是可以动态变化的

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p\n", &slice[0])
}
