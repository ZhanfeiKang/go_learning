package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式1
	// var p1 Person
	// fmt.Println(p1)

	// 方式2
	p2 := Person{}
	p2.Name = "tom"
	p2.Age = 18
	fmt.Println("p2 :", p2)

	p3 := Person{"mary", 20}
	fmt.Println("p3 :", p3)

	// 方式3
	var p4 *Person = new(Person)
	// 因为p4是一个指针，因此标准的给字段赋值方式是：
	// (*p4).Name = "smith"  也可以这样写  p4.Name = "smith",
	// 设计者为了程序员方便，会在底层进行处理 会给 p3 加上取值运算
	(*p4).Name = "smith"
	p4.Age = 30
	fmt.Println("p4 :", *p4)

	// 方式4
	var p5 *Person = &Person{}
	// 因为person是一个指针,原理同上
	// var p5 *Person = &Person{"kkite", 22}
	(*p5).Name = "kkite"
	p5.Age = 22
	fmt.Println("p5 :", *p5)
}
