package main

import "fmt"

func main() {
	var name string
	var pwd string
	var loginChance = 3

	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)

		if name == "张无忌" && pwd == "888" {
			fmt.Println("恭喜你登录成功")
			break
		} else {
			loginChance--
			fmt.Printf("你还有%v次机会，请珍惜\n", loginChance)
		}
	}

	if loginChance == 0 {
		fmt.Println("机会用完，没有登录成功！")
	}
}
