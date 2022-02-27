package main

import "fmt"

func main() {
	// 使用for range 遍历map
	cities := make(map[string]string)

	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	fmt.Println(cities)

	for k, v := range cities {
		fmt.Println(k, ":", v)
	}

	fmt.Printf("cities有%v对key-value", len(cities))

	// 遍历多重map
	stuMap := make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 2)
	stuMap["stu01"]["name"] = "tom"
	stuMap["stu01"]["sex"] = "男"
	stuMap["stu01"]["address"] = "上海"

	stuMap["stu02"] = make(map[string]string, 2)
	stuMap["stu02"]["name"] = "mary"
	stuMap["stu02"]["sex"] = "女"
	stuMap["stu02"]["address"] = "云梦"

	for k1, v1 := range stuMap {
		fmt.Println(k1, ":")
		for k2, v2 := range v1 {
			fmt.Printf("\t %v : %v\n", k2, v2)
		}
	}
}
