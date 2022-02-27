package main

import "fmt"

func main() {
	// string的基本使用
	var address string = "北京长城 110 hello world！"

	fmt.Println(address)

	// 字符串一旦赋值了，字符串就不能修改了：在Go中字符串是不可变的
	// var str = "hello"
	// str[0] = 'a' // 这里就不能去修改str的内容，即go中的字符串是不可改变的

	// 使用反引号 ``
	str2 := `abc\nabc`
	str3 := `
	func main() {
		// string的基本使用
		var address string = "北京长城 110 hello world！"
	
		fmt.Println(address)
	
		// 字符串一旦赋值了，字符串就不能修改了：在Go中字符串是不可变的
		// var str = "hello"
		// str[0] = 'a' // 这里就不能去修改str的内容，即go中的字符串是不可改变的
	
		// 使用反引号 
		str2 := abc\nabc
		str3 := 
		fmt.Println(str2)
	}`
	fmt.Println(str2, str3)

	// 字符串拼接方式
	var str4 = "hello" + "world!"
	str4 += " haha"

	fmt.Println(str4)

	// 当一个字符串太长时间，怎么办，可以分行写, + 号需要保留在上一行
	var str5 = "hello " + "world " + "hello " + "world " +
		"hello " + "world " + "hello " + "world " +
		"hello " + "world "

	fmt.Println(str5)
}
