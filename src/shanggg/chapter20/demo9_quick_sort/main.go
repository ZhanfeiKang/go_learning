package main

import (
	"fmt"
	"time"
)

// 快速排序
// 说明
// 1. left 表示 数组左边的下标
// 2. right 表示 数组右边的下标
// 3. array 表示 要排序的数组
func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right

	// pivot 是中轴，支点
	pivot := array[(left+right)/2]
	temp := 0

	// for 循环的目标是将比 pivot 小的数放到 左边
	// 					比 pivot 大的数放到 右边
	for l < r {

		// 从pivot的 左边 找到 >=pivot 的值
		for array[l] < pivot {
			l++
		}

		// 从pivot的 右边 找到 <=pivot 的值
		for array[r] > pivot {
			r--
		}
		// l>=r : 表明本次分解任务完成，break
		if l >= r {
			break
		}

		// 交换
		temp = array[l]

		array[l] = array[r]
		array[r] = temp

		// 特殊处理，不然会死循环
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// 如果 l==r，再移动下,中间的pivot不参与重复排序
	if l == r {
		l++
		r--
	}

	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

func fastSort(start int, end int, arr *[6]int) {
	if start >= end {
		return
	}
	left := start
	right := end
	node := arr[left]
	// 找出所有比node小的值，放在node前面，
	// 所有比node大的值放在node后面
	for left < right {

		for left < right && arr[right] >= node {
			right--
		}
		arr[left] = arr[right]

		for left < right && arr[left] <= node {
			left++
		}
		arr[right] = arr[left]
	}

	arr[left] = node

	fastSort(start, left-1, arr)
	fastSort(left+1, end, arr)

}

func main() {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	// QuickSort(0, len(arr)-1, &arr)
	start := time.Now().UnixNano()
	fastSort(0, len(arr)-1, &arr)
	end := time.Now().UnixNano()
	fmt.Println("arr: ", arr)

	fmt.Printf("time: %v", end-start)

	// {-9, 78, 0, 23, -567, 70}
	// {-9, -567, 0, 23, 78, 70}

	// 左边{-567, -9, 0, 23, 78, 70}
}
