package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	for num := 1; num <= 80000; num++ {
		flag := true
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该 num 不是素数
				flag = false
				break
			}
		}

		if flag {
			// 将 num 放入 primeChan
			// primeChan <- num
		}
	}
	end := time.Now().Unix()
	fmt.Println("普通方法耗时:", end-start)
}
