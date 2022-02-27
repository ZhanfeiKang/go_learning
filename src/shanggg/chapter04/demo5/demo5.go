package main

import "fmt"

func main() {
	// 位运算的演示
	fmt.Println(2 & 3)  // 2
	fmt.Println(2 | 3)  // 3
	fmt.Println(2 ^ 3)  // 3
	fmt.Println(-2 ^ 2) // -4

	fmt.Println(1 >> 2) // 0000 0000   0
	fmt.Println(1 << 2) // 0000 0100   4

}
