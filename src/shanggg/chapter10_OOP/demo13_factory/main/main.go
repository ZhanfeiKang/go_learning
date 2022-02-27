package main

import (
	"fmt"
	"shanggg/chapter10_OOP/demo13_factory/model"
)

func main() {
	// 创建一个Student的实例
	// var stu = model.Student{
	// 	Name:  "tom",
	// 	Score: 78.5,
	// }

	// 当student结构体首字母是小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom~", 88.8) // stu是个指针

	fmt.Println(*stu)
	fmt.Println("name :", stu.Name)
	fmt.Println("score :", stu.GetScore()) // 通过stu.GetScore()获取score
}
