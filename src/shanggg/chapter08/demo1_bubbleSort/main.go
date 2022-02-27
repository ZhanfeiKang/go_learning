package main

import (
	"fmt"
)

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前 :", *arr)
	temp := 0 // 临时变量(用作于交换)
	for i := 0; i < len(*arr); i++ {

		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] { // > 从小到大
				// 交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}

		fmt.Printf("第%d次排序后arr :%v\n", i+1, *arr)

	}
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}

	BubbleSort(&arr)
	fmt.Println("排序完成！\n arr :", arr)
}
