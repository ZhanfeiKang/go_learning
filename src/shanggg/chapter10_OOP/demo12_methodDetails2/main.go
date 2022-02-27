package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

// 编写一个方法，可以改变i的值

func main() {

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(stu)
	// 如果你实现了 *Student 类型的 String 方法，就会自动调用 （String返回值必须是string类型）
	fmt.Println(&stu)
}
