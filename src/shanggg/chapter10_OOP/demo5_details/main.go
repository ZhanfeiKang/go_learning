package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	// r1有四个int，在内存中是连续分布
	// 打印地址
	fmt.Printf("r1.leftUp.x 地址: %p \nr1.leftUp.y 地址: %p\n"+
		"r1.rightDown.x 地址: %p \nr1.rightDown.y 地址: %p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// r2有两个 *Point类型，这两个 *Point类型本身地址也是连续的，
	// 但是他们指向的地址不一定是连续的
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}
	fmt.Printf("r2.leftUp 本身地址: %p \nr2.rightDown 本身地址: %p\n",
		&r2.leftUp, &r2.rightDown)

	// 他们指向的地址不一定是连续... 这个要看系统运行时内存是如何分配的
	fmt.Printf("r2.leftUp 指向地址: %p \nr2.rightDown 指向地址: %p\n",
		r2.leftUp, r2.rightDown)

}
