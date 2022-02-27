package main

import "fmt"

func main() {
	// 某人有100,000元，每经过一次路口，需要交费，规则如下
	// 当现金>50000时，每次交5%
	// 当现金<=50000时，每次交1000
	var (
		cash  float64 = 100000
		count int     = 0
	)

	for {
		if cash > 50000 {
			cash -= cash * 0.05
			count++
		} else if cash <= 50000 && cash >= 1000 {
			cash -= 1000
			count++
		} else {
			break
		}
	}

	fmt.Printf("该人可以经过%v次路口，最后还剩%v元", count, cash)
}
