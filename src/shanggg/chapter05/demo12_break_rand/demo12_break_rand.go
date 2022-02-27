package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 我们为了生成一个随机数，还需要个rand设置一个种子
	// time.Now().Unix(): 返回一个从1970年01月01日00:00:00到现在的一个秒数
	rand.Seed(time.Now().Unix())
	// 如何随机的生成一个1-100的整数
	n := rand.Intn(100) + 1 // [0,100)
	fmt.Println(n)

	// 编写一个无限循环的控制，然后不停的随机生成数，当生成了99时，就退出这个无限循环 -> break
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano()) // 种子使用纳秒，更加随机
		n := rand.Intn(100) + 1
		fmt.Println(n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共使用了", count)
}
