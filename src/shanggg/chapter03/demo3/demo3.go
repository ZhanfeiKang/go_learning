package main

import "fmt"

// 定义全局变量
var n1 = 100
var n3 = 200
var name = "jack"

// 上面的声明方式，也可以改成一次性声明
var (
	n4    = 300
	n5    = 900
	name2 = "mary"
)

func main() {
	// 如何一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 一次性声明多个变量的方式2
	// var n1, name, n3 = 100, "tom", 888
	// fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	// 方式3
	// n1, name, n3 := 100, "tom", 888
	fmt.Println("n1=", n1, "name=", name, "n3=", n3)
	fmt.Println("n4=", n1, "name2=", name2, "n5=", n5)
}
