package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Age = 10
	p1.Name = "小明"

	var p2 *Person = &p1

	fmt.Println((*p2).Name)
	fmt.Println(p2.Age)

	p2.Name = "tom~"
	fmt.Printf("p2.Name=%v p1.Name=%v \n", p2.Name, p1.Name)    // tom~  tom~
	fmt.Printf("p2.Name=%v p1.Name=%v \n", (*p2).Name, p1.Name) // tom~  tom~

	fmt.Printf("p1的地址：%p\n", &p1)
	fmt.Printf("p2的地址:%p, p2的值:%p", &p2, p2)

	// *p2.Name = "tom"  不能这样写，因为   . 的优先级  高于   * 的优先级
}
