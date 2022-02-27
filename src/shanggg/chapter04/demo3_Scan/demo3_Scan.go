package main

import "fmt"

func main() {
	// 方式1：scanln

	var (
		name   string
		age    byte
		sal    float32
		isPass bool
	)

	// fmt.Println("请输入姓名：") // 输入后值会自动转成变量声明的类型
	// fmt.Scanln(&name)     // 关键点 ： 传地址

	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入是否通过考试：")
	// fmt.Scanln(&isPass)

	// fmt.Printf("名字是：%v \n年龄是：%v \n薪水是：%v\n是否通过考试：%v\n", name, age, sal, isPass)

	// 方式2：Scanf
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字是：%v \n年龄是：%v \n薪水是：%v\n是否通过考试：%v\n", name, age, sal, isPass)

}
