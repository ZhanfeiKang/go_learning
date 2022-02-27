package main

import "fmt"

/*
	斐波那契数
	使用递归，求出斐波那契数 1,1,2,3,5,8,13...
	1) 当n=1||n=2，返回1
	2) 当n>=2, 返回前两个数的和 f(n-1)+f(n-2)
*/

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {

	var n int = 1
	fmt.Print("请输入n：")
	fmt.Scanln(&n)
	fmt.Println(fbn(n))

}
