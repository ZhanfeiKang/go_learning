package main

import "fmt"

func main() {
	//1. for
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}
	fmt.Println()

	//2. for  range
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	slice2 := slice[1:2]
	slice2[0] = 100 // slice和arr对应地址都会改变

	fmt.Println("slice2 :", slice2)
	fmt.Println("slice :", slice)
	fmt.Println("arr :", arr)
}
