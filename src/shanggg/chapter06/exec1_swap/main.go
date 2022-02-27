package main

import "fmt"

func swap(n1 *int, n2 *int) {
	// 定义一个临时变量
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Printf("a=%v,b=%v\n", a, b)
}
