package main

import "fmt"

type Ainterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Student Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i=", i)
}

type Binterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say~")
}

func main() {
	var stu Stu
	// var a Ainterface //不能实例化接口
	var a Ainterface = stu
	a.Say()

	var i integer = 10
	var b Ainterface = i
	b.Say() //integer Say i= 10

	// Monster实现了Ainterface和Binterface两个接口
	var monster Monster
	var a2 Ainterface = monster
	var b2 Binterface = monster

	a2.Say()
	b2.Hello()

	var t interface{} = monster // 空接口，可以把任何一个变量赋给空接口
	fmt.Println(t)
}
