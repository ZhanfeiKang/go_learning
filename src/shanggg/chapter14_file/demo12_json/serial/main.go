package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"` // 反射机制
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	// 演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "1555-5-5",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	// 将monster进行序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败 err: ", err)
		return
	}
	// 输出序列化后的结果
	fmt.Println("monster序列化后: ", string(data))
}

// 将map进行序列化
func testMap() {
	// 定义一个map
	// var a map[string]interface{}
	// 使用map，需要make
	a := make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["skill"] = "火云洞"

	// 将a这个map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败 err: ", err)
		return
	}
	// 输出序列化后的结果
	fmt.Println("a map序列化后: ", string(data))
}

// 对切片进行序列化,切片 []map[string]interface{}
func testSlice() {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 20
	m2["address"] = [2]string{"墨西哥", "芝加哥"}
	slice = append(slice, m2)

	// 将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败 err: ", err)
		return
	}
	// 输出序列化后的结果
	fmt.Println("slice序列化后: ", string(data))
}

// 对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64() {
	// 对num1序列化
	var num1 float64 = 2345.567
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化失败 err: ", err)
		return
	}
	// 输出序列化后的结果
	fmt.Println("num1序列化后: ", string(data))
}

func main() {
	// 将结构体，map，切片进行序列化
	testStruct()
	testMap() // map序列化后，是无序的
	testSlice()
	testFloat64()
}
