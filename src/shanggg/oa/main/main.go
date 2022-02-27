package main

import (
	"fmt"
	util "shanggg/oa/utils"
)

func main() {
	var (
		n1      float64 = 1.2
		n2      float64 = 2.3
		operate byte    = '*'
	)

	result := util.Cal(n1, n2, operate)
	fmt.Printf("result=%.2f", result)
}
