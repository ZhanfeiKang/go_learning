package main

import "fmt"

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(rune(36)) // -> $
		fmt.Println("str=", str)
		return n
	}
}
func main() {
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16
}
