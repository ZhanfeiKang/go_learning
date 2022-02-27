package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("你好 瑶瑶~")
	}

	// 第二种
	j := 1
	for j <= 10 {
		fmt.Println("你好 瑶瑶~")
		j++
	}

	// 第三种：死循环，配合break
	k := 1
	for {
		if k <= 10 {
			fmt.Println("你好 瑶瑶~")
		} else {
			break
		}
		k++
	}
}
