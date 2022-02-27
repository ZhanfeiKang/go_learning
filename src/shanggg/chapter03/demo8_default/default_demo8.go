package main

import "fmt"

func main() {
	var a int
	var b float32
	var c float64
	var isMarryied bool
	var name string

	// 这里的%v 表示按照变量的值输出
	fmt.Printf("a=%d, b=%f, c=%f, isMarrayied=%v, name=%v", a, b, c, isMarryied, name)
}
