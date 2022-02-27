package main

import "fmt"

func sum(n1 int, n2 int) int {

	// 当执行到defer是，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
	// 当函数执行完后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok n1=", n1) // defer	3. ok3 n1=20
	defer fmt.Println("ok n2=", n2) // defer    2. ok2 n2=10

	n1++ // n1=21
	n2++ // n2=11

	res := n1 + n2
	fmt.Println("ok3 res=", res) // 1. ok3

	return res
}

func main() {
	res := sum(20, 10)
	fmt.Println("res=", res) // 4.
}
