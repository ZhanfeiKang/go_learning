package main

import (
	"fmt"
	"reflect"
)

// 通过反射，修改
// num int 的值
// 修改 student 的值

func reflect01(b interface{}) {

	rVal := reflect.ValueOf(b)
	// 看看 rVal 的kind是
	fmt.Println("rVal kind: ", rVal.Kind()) // ptr

	// rVal.SetInt(20) // 错误

	// rVal.Elem() 返回v指针持有的值
	// Elem返回v持有的接口保管的值的Value的封装，或者v持有的指针指向的值的Value封装
	rVal.Elem().SetInt(20)
}

func main() {

	num := 10
	reflect01(&num)

	fmt.Println(num)
}
