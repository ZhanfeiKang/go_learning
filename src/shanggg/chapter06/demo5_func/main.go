package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

type myFunType func(int, int) int

// 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
// func myFun(funvar func(int, int) int, num1 int, num2 int) int {
// 	return funvar(num1, num2)
// }
func myFun(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSumAndSub(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	a := getSum
	fmt.Printf("a的类型%T, getSum类型是%T\n", a, getSum)

	res := a(10, 40)
	fmt.Println("res=", res)

	res2 := myFun(a, 20, 21)
	fmt.Println("res2=", res2)

	a1, b1 := getSumAndSub(14, 5)
	fmt.Println("a1=", a1, "b1=", b1)
}
