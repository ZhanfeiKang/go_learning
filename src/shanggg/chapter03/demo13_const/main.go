package main

import "fmt"

// 常量仍然通过首字母的大小写来控制常量的访问范围

func main() {
	var num int

	// 常量在声明的时候必须给值
	const tax int = 0
	// 常量是不能修改的
	// tax = 10

	fmt.Println(num, tax)

	// 常量只能修饰bool、数值类型(int,float系列)、string类型

	const num2 = 9 / 3
	fmt.Println("num2: ", num2)

	const (
		a = iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d) // 0,1,2,3
}
