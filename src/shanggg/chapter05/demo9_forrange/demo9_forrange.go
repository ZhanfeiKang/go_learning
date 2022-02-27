package main

import "fmt"

func main() {
	// 字符串遍历方式-传统方式
	var str string = "hello 北京~"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}

	// for range
	str = "kkite l 瑶瑶"
	fmt.Println()
	for index, val := range str {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}
}
