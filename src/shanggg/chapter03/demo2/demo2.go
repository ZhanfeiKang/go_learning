package main

import "fmt"

func main() {
	// 定义变量 声明变量
	// 第一种
	// int 默认值为0
	var i int
	fmt.Println("i=", i)

	// 第二种：根据值自行判定变量类型
	var num = 10.11
	fmt.Println("num=", num)

	// 第三种：省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	// 下面的方式等价于 var name string 	name = "tom"
	name := "tom"
	fmt.Println("name=", name)
}
