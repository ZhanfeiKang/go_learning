package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成5个数，并将其反转打印
// 反转打印，倒数第一个和第一个交换，交换次数是 len/2

func main() {
	var intArr [5]int
	var length = len(intArr)
	// 为了每次生成的随机数不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		intArr[i] = rand.Intn(100) //  [0,100)
	}
	fmt.Println("交换前：", intArr)

	// 反转打印
	temp := 0 // 做一个临时变量
	for i := 0; i < length/2; i++ {
		temp = intArr[length-1-i]
		intArr[length-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交换后：", intArr)
}
