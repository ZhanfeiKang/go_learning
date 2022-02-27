package main

import (
	"fmt"
	"shanggg/chapter11_OOp2/exer1_Account/model"
)

func main() {
	account := model.NewAccount("gs88", "666666", 40)
	if account != nil {
		fmt.Println("创建成功 :", *account)
	} else {
		fmt.Println("创建失败")
	}

}
