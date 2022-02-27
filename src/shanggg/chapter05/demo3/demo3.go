package main

import "fmt"

func main() {
	var n1 int32 = 10
	var n2 int32 = 5
	var m int32 = n1 + n2

	if m%3 == 0 {
		if m%5 == 0 {
			fmt.Println("能被3又能被5整除")
		} else {
			fmt.Println("能被3不能被5整除")
		}
	} else {
		if m%5 == 0 {
			fmt.Println("不能被3，但能被5整除")
		} else {
			fmt.Println("不能被3也不能被5整除")
		}
	}
}
