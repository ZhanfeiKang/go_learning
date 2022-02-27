package main

import (
	"fmt"
)

func main() {
	// map的声明和注意事项
	// var a map[string]string
	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	var a map[string]string = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松" // key不能重复，覆盖
	a["no3"] = "吴用" // value可以重复
	fmt.Println(a)

	// 第二种使用方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 第三种方式
	heros := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
		"hero3": "鲁智深",
	}
	fmt.Println(heros)

	// 案例
	/*
		我们要存放3个学生信息，每个学生有name和sex信息
		思路：map[string]map[string]string
	*/
	stuMap := make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 2)
	stuMap["stu01"]["name"] = "tom"
	stuMap["stu01"]["sex"] = "男"

	stuMap["stu02"] = make(map[string]string, 2)
	stuMap["stu02"]["name"] = "mary"
	stuMap["stu02"]["sex"] = "女"

	fmt.Println(stuMap)
}
