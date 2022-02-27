package main

import "fmt"

func main() {
label2:
	for i := 0; i < 4; i++ {
		// label1: // 设置一个标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break 会默认跳出最近的一个for循环
				// break label1
				break label2
			}
			fmt.Println("j=", j)
		}
	}
}
