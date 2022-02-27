package utils

import "fmt"

// 为了让其他包的文件使用Cal函数，需要将c大写，类似其他语言的public
func Cal(n1 float64, n2 float64, operate byte) float64 {
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
