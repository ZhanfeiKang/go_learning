package main

import "fmt"

func main() {
	cities := make(map[string]string)

	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	fmt.Println(cities)

	// 因为no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)

	// 演示删除
	delete(cities, "no1")
	fmt.Println(cities)

	delete(cities, "no100")
	fmt.Println(cities)

	// map查找
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有 no1 ，key值为%v\n", val)
	} else {
		fmt.Println("没有 no1 key")
	}

	//如果希望一次性删除所有的key
	//1.遍历，逐一删除
	//2.直接make一个新的空间   推荐***
	cities = make(map[string]string)
	fmt.Println(cities)
}
