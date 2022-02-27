package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你年龄大于18，要对自己的行为负责")
	} else {
		fmt.Println("你的年龄不大，这次放过你了")
	}
}
