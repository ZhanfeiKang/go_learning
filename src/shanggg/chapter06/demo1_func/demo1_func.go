package main

import "fmt"

func cal(n1 float64, n2 float64, operate byte) float64 {
	var res float64

	switch operate {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("输入符号错误...")
	}
	return res
}

func main() {
	var (
		n1      float64 = 1.2
		n2      float64 = 2.3
		operate byte    = '+'
	)

	result := cal(n1, n2, operate)
	fmt.Println("result=", result)

}
