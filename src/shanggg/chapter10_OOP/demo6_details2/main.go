package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {
	var a A
	var b B
	// a = b	错
	a = A(b) // 需要有完全相同的字段、个数、类型
	fmt.Println(a, b)
}
