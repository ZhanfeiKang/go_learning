package main

import "fmt"

// + 号的使用
func main() {
	var i, j = 1, 2
	var r = i + j
	fmt.Println("r=", r)

	var str1 = "yaoyao "
	var str2 = "hao ke ai"
	var str3 = str1 + str2
	fmt.Println(str3)
}
