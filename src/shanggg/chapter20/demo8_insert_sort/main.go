package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	插入排序基本思想:
	1.把n个待排序的元素看为一个有序表和一个无序表，
	2.开始时，有序表中只包含一个元素，无需表中包含有n-1个元素，
	3.排序过程中每次从无序表中取出第一个元素，把它的排序码依次与有序表元素的排序码进行比较，将它插入到有序表中的适当位置，使之成为新的有序表。
*/

func InsertSort(arr *[80000]int) {

	for i := 1; i < len(arr); i++ {
		// 完成第一次，给第二个元素找到合适的位置并插入
		insertVal := arr[i]
		insertIndex := i - 1

		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal { // 往前走
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != i { // 本身就在原来的位置
			arr[insertIndex+1] = insertVal
		}
		// fmt.Printf("第%d次插入后的结果: %v\n", i, *arr)
	}

	/*

		// 完成第2次，给第二个元素找到合适的位置并插入
		insertVal = arr[2]
		insertIndex = 2 - 1

		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal { // 往前走
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != 2 { // 本身就在原来的位置
			arr[insertIndex+1] = insertVal
		}
		fmt.Println("第2次插入后的结果: ", *arr)

	*/
}

func main() {

	// arr := [7]int{23, 0, 12, 56, 34, -1, 55}
	// fmt.Println("\t原始数组 : ", arr)

	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}

	start := time.Now().Unix()
	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("插入排序耗时：%d秒", end-start)
	// 插入排序耗时：2秒

}
