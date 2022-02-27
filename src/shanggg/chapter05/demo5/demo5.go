package main

import (
	"fmt"
	"math"
)

func main() {
	// 求ax^2+bx+c=0 方程的根。a,b,c分别为函数的参数，如果：b^2-4ac>0,则有两个解；
	// b^2-4ac=0,则有一个解；b^2-4ac<0，则无解；
	// 提示1：x1 = (-b+sqrt(b^2-4ac))/2a;   x2 = (-b-sqrt(b^2-4ac))/2a
	// 提示2：math.Sqrt(num)；可以求平方根 需要引入math包

	var (
		a  float64 = 3.0
		b  float64 = 100.0
		c  float64 = 6.0
		x1 float64
		x2 float64
	)

	m := b*b - 4*a*c

	if m > 0 {
		x1 = (-b + math.Sqrt(m)) / (2 * a)
		x2 = (-b - math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v x2=%v\n", x1, x2)
	} else if m == 0 {
		x1 = (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=x2=%v\n", x1)
	} else {
		fmt.Println("无解")
	}
}
