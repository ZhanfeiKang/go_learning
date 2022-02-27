package main

import "fmt"

func main() {
	var intAir [3]int32 // int32每个元素占4个字节
	fmt.Printf("intAir的地址：%p\n", &intAir)
	fmt.Printf("intAir[0]的地址：%p\n", &intAir[0])
	fmt.Printf("intAir[1]的地址：%p\n", &intAir[1])

	// 4种初始化数组的方式
	// 1.
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01 =", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02 =", numArr02)

	// 这里的 [...] 是规定的写法
	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03 =", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04 =", numArr04)
}
