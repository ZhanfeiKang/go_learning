package main

import "fmt"

// 可变参数与的使用
func sum(n1 int, args ...int) int {
	sum := n1
	// 遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i] // args[0] 表示取出args切片的第一个元素值，其他依此类推
	}
	return sum
}

func main() {
	res := sum(10, 0, -1, 90, 10)
	fmt.Println("res=", res)
}
