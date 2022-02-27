package main

import "fmt"

func main() {
	// string底层是一个byte数组，所以也可以进行切片处理
	str := "hello@kkite"
	// 使用切片获取到kkite
	slice := str[6:]
	fmt.Println("slice :", slice)

	// string是不可变的
	// str[0] = 'z' 错误操作
	arr1 := []byte(str) // 转成byte切片
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str :", str)

	// 细节，我们转成[]byte后，可以处理英文和数字，但不能处理中文
	// 解决方法，将string 转成 []rune 即可，因为 []rune 是按字符处理，兼容汉子
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str :", str)
}
