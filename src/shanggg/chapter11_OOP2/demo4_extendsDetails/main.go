package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A say ok", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
}

func (b *B) SayOk() {
	fmt.Println("B hello", b.Name)
}

func main() {
	var b B

	b.A.Name = "tom"
	b.A.age = 19
	b.A.SayOk()
	b.A.hello()

	// 可以简化成下面的写法
	var b1 B
	b1.Name = "mary"
	b1.age = 22
	b1.SayOk()
	b1.hello()
}
