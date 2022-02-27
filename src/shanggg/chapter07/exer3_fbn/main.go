package main

import "fmt"

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	// 1,1,...
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	/*
		1. 声明一个函数 fbn(n ,int)  ([]uint64)
		2. 编程fbn(n int) 进行for循环来存放斐波那契数列   0->1, 1->1
	*/
	fbnSlice := fbn(10)
	fmt.Println("fnbSlice :", fbnSlice)
}
