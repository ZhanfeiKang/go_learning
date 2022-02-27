package main

import "fmt"

// 切片的创建
func main() {
	// 1.
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	slice := intArr[1:3]
	fmt.Println(slice)

	// 2.通过make创建
	var slice2 []int = make([]int, 5, 10)
	slice2[1] = 10
	slice2[3] = 20
	// 对于切片，必须make后才能使用
	fmt.Println("slice2 :", slice2)
	fmt.Println("slice2的size :", len(slice2))
	fmt.Println("slice2的cap :", cap(slice2))

	// 3.
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strString :", strSlice)
	fmt.Println("strString size :", len(strSlice))
	fmt.Println("strString cap :", cap(strSlice))
}
