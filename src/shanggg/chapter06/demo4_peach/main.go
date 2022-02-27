package main

import "fmt"

/*
	有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个；以后猴子每天都吃其中的一半，然后再多吃一个。
	当到第十天时，想再吃时（还没吃），发现只有1个桃子了
	问：最初共多少个桃子？
*/

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
		return 0 // 表示没有得到正确数量
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func main() {
	fmt.Println("第一天桃子树量是=", peach(1))
}
