package main

import (
	"fmt"
)

func yuejieError() {
	defer func() {
		err := recover() //会捕获一个错误
		if err != nil {
			fmt.Println("123")
		}
	}()
	var arr1 [3]int
	for i := 0; i <= 3; i++ {
		arr1[i] = 1
	}
	// var age int
	// age = sdasda

}
func main() {
	yuejieError()
}
