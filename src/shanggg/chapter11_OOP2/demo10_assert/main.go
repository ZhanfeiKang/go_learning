package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point // ok
	// 如何将 a 赋给一个Point变量
	var b Point
	// b = a // error
	b = a.(Point) // 类型断言
	fmt.Println(b)

	// 类型断言的其他案例
	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2 // 空接口可以接收任意类型
	// y := x.(float32)
	// fmt.Printf("y 的类型是: %T,值是: %v", y, y)

	// 带检测的类型断言
	var x interface{}
	var b2 float32 = 1.1
	x = b2 // 空接口可以接收任意类型
	y, ok := x.(float32)
	if ok {
		fmt.Println("convert success")
		fmt.Printf("y 的类型是: %T,值是: %v\n", y, y)
	} else {
		fmt.Println("转换失败...")
	}

	fmt.Println("继续执行~~")
}
