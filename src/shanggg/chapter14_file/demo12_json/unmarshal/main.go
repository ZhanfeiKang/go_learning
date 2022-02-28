package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

// 演示将json字符串，反序列化成结构体
func unmarshalStruct() {
	// str在项目开发中，是通过网络传输获取到..或者是读取文件获取到
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"1555-5-5\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	// 定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmarshal err: ", err)
		return
	}
	fmt.Println("反序列化后 monster: ", monster)
}

func unmarshalMap() {
	str := "{\"age\":30,\"name\":\"红孩儿\",\"skill\":\"火云洞\"}"

	// 定义一个map
	// 注意：反序列化map，不需要make，因为make操作已经被封装到 Unmarshal函数中了
	var a_map map[string]interface{}

	err := json.Unmarshal([]byte(str), &a_map)
	if err != nil {
		fmt.Println("unmarshal err: ", err)
		return
	}
	fmt.Println("反序列化后 a_map: ", a_map)
}

// 演示将json字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":7,\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"芝加哥\"],\"age\":20,\"name\":\"tom\"}]"

	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("unmarshal err: ", err)
		return
	}
	fmt.Println("反序列化后 slice: ", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
