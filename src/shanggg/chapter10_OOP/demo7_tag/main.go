package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}

	// 2.将monster变量序列化为json格式字符串
	// json.Marshal 函数中使用反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误:", err)
	}
	fmt.Println("jsonStr :", string(jsonStr))
	fmt.Println(monster)
}
