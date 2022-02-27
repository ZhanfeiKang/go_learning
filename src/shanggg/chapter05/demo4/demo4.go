package main

import "fmt"

func main() {
	var year int = 2020

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("是润年~")
	}
}
