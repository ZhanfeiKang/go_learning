package main

import "fmt"

func main() {
	// 打印金字塔
	/*
		  *			1层 1个*	规律：2 * 层数 - 1	| 空格 2	规律：总层数 - 当前层数
		 ***		2层 3个*	规律：2 * 层数 - 1	| 空格 1	规律：总层数 - 当前层数
		*****		3层 5个*	规律：2 * 层数 - 1	| 空格 0	规律：总层数 - 当前层数
	*/

	var totallevel int = 15

	// i表示层数
	for i := 1; i <= totallevel; i++ {
		// 在打印*前先打印空格
		for k := 1; k <= totallevel-i; k++ {
			fmt.Print(" ")
		}

		// j 表示每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i <= totallevel; i++ {
		// 在打印*前先打印空格
		for k := 1; k <= totallevel-i; k++ {
			fmt.Print(" ")
		}

		// j 表示每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totallevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
