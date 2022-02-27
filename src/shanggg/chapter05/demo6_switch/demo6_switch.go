package main

import "fmt"

func test(b byte) byte {
	return b + 1
}

func main() {
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	switch test(key) {
	case 'a':
		fmt.Println("周一，猴子穿新衣")
	case 'b':
		fmt.Println("周二，猴子当小二")
	case 'c':
		fmt.Println("周三，猴子爬雪山")
	//...
	default:
		fmt.Println("输入有误...")
	}
}
