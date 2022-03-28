package main

import "fmt"

type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Println("我是回调，x: ", x)
	return x
}

func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Println("我是回调,x: ", x)
		return x
	})
}
