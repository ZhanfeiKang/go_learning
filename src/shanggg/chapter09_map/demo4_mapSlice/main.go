package main

import "fmt"

func main() {
	// 演示map切片的使用
	/*
		要求：使用一个map来记录monster的信息 name 和 age，也就是说一个
		monster对应一个map，并且妖怪的个数可以动态的增加=>map切片
	*/
	// var monsters []map[string]string
	var monsters = make([]map[string]string, 2) // slice 的make，准备放入两个妖怪
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "100"
	}

	// 使用切片的动态增长
	newMonster := map[string]string{
		"name": "火云邪神",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
}
