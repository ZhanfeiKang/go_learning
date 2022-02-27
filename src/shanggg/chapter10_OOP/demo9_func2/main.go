package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Println(p.Name + "是一个goodman~")
}

func (p Person) jisuan() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(p.Name+"计算的结果是：", sum)
}

func (p Person) jisuan2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(p.Name+"计算的结果是：", sum)
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var tom Person
	tom.Name = "tom"
	tom.speak()
	tom.jisuan()
	tom.jisuan2(10)

	res := tom.getSum(10, 20)
	fmt.Println("res :", res)
}
