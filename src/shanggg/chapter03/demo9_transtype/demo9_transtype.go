package main

import "fmt"

// 基本数据类型的转换
func main() {
	var i int32 = 100

	// 希望将 i => float32
	var n1 float32 = float32(i)

	var n2 int8 = int8(i)

	fmt.Printf("i=%v, n1=%v, n2=%v", i, n1, n2)
}
