package main

import "fmt"

var (
	// 全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {

	// 匿名函数
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1=", res1)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(5, 2)
	res3 := a(5, 8)
	fmt.Println("res2=", res2)
	fmt.Println("res3=", res3)

	res4 := Fun1(5, 5)
	fmt.Println("res4=", res4)
}
