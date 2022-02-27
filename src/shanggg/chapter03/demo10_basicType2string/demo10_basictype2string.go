package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型转string
func main() {
	var (
		num1   int     = 99
		num2   float64 = 23.465
		b      bool    = true
		myChar byte    = 'h'
		str    string
	)

	// 第一种

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b) // %t 布尔值
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar) // %t 布尔值
	fmt.Printf("str type %T str=%q\n", str, str)

	// 第二种：使用strconv包中的函数
	var (
		num3 int     = 99
		num4 float64 = 23.456
		b2   bool    = true
	)

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	// 说明：'f'格式 10：表示小数位保留十位   64：表示这个小数是float64
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	var num5 int = 4567
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str)
}
