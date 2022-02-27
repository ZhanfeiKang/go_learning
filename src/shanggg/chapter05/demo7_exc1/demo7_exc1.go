package main

import "fmt"

func main() {
	var letter byte
	fmt.Println("请输入小写字母：")
	fmt.Scanf("%c", &letter)

	switch letter {
	case 'a', 'b', 'c', 'd', 'e':
		letter -= 32
		fmt.Printf("%c", letter)
	default:
		fmt.Println("other...")
	}
}
