package main

import "fmt"

func main() {
	// 1.年龄变量
	// 2.从控制台接收年龄
	// 3.if判断
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你年龄大于18，要对自己的行为负责")
	}

	// golang允许在if中，直接定义一个变量
	// if age := 20; age > 18 {
	// 	fmt.Println("你年龄大于18，要对自己的行为负责")
	// }
}
