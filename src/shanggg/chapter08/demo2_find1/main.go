package main

import "fmt"

func main() {
	/*
		1. 定义一个数组，白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王 字符串数组
		2. 从控制台接收一个名字，一次比较，如果发现有，提示
	*/

	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名:")
	fmt.Scanln(&heroName)

	// 顺序查找：第一种方式
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("找到%v，下标%v\n", heroName, i)
			break
		} else if i == len(names)-1 {
			fmt.Printf("没有找到%v\n", heroName)
		}
	}

	// 顺序查找：方式2
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v，下标%v\n", heroName, index)
	} else {
		fmt.Printf("没有找到%v\n", heroName)
	}
}
