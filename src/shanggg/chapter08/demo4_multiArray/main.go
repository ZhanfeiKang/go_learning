package main

import "fmt"

func main() {
	// 二维数组
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/

	var arr [4][6]int

	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	// 遍历二维数组
	fmt.Println("双for遍历：")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("\nfor...range...遍历：")
	for _, v := range arr {
		for _, v2 := range v {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1])

	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])

	// 创建方式二
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3 :", arr3)
}
