package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
}

type D struct {
	a A // 组合关系，不叫继承了
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	var c C
	// c.Name = "tom"	// 会报错
	c.B.Name = "tom"

	var d D
	// d.Name = "mary" // 会报错
	d.a.Name = "mary"

	tv1 := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}}
	tv2 := TV{
		Goods{
			Name:  "电视机002",
			Price: 6999.0,
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}

	tv3 := TV2{&Goods{"电视机003", 5000.99}, &Brand{"创维", "河南"}}
	tv4 := TV2{
		&Goods{
			Name:  "电视机004",
			Price: 6999.0,
		},
		&Brand{
			Name:    "长虹",
			Address: "四川",
		},
	}

	fmt.Println("tv1 :", tv1)
	fmt.Println("tv2 :", tv2)
	fmt.Println("tv3 :", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4 :", *tv4.Goods, *tv4.Brand)

}
