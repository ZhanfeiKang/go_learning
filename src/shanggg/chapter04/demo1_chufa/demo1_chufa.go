package main

import "fmt"

func main() {
	// 如果参与运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4 // 2
	fmt.Println(n1)

	var n2 float32 = 10.0 / 4
	fmt.Println(n2)
}
