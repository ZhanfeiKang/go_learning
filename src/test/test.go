package main

import "fmt"

func main() {
	var (
		// n1       float64
		// n2       float64
		operator byte
	)
	fmt.Println("请输入：")
	fmt.Scanf("%d\n", &operator)
	new_o := fmt.Sprintf("%c", operator)
	fmt.Println(new_o)
}
