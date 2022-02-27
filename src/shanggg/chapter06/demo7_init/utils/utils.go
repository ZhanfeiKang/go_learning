package utils

import "fmt"

var Age int
var Name string

// Age 和 Name全局变量，需要在main.go中使用
// 但是我们需要初始化Age和Name

// 让init函数完成初始化工作
func init() {
	Age = 100
	Name = "tom"

	fmt.Println("utils中的init()...")
}
