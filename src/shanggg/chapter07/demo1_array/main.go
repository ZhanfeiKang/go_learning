package main

import "fmt"

func main() {
	// 1.定义一个数组
	var hens [6]float64
	// 2.赋值
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	// 3.遍历数组
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	// 4.求总
	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("totalWeight=%v, avgWeight=%v", totalWeight, avgWeight)
}
