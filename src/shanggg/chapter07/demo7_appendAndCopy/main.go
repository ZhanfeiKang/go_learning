package main

import "fmt"

func main() {
	// 用append内置函数，可以对切片进行动态追加
	var slice []int = []int{100, 200, 300}
	fmt.Println("slice :", slice)
	slice = append(slice, 400, 500, 600) // 需要用一个切片接收,重新赋值给slice
	fmt.Println("slice :", slice)

	slice1 := append(slice, 123)
	slice = append(slice, slice1...) // 可以追加一个切片
	fmt.Println("slice :", slice)

	// 切片拷贝
	var slice3 []int = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, 10)
	copy(slice4, slice3)
	fmt.Println("slice3 :", slice3) // [1 2 3 4 5]
	fmt.Println("slice4 :", slice4) // [1 2 3 4 5 0 0 0 0 0]
}
