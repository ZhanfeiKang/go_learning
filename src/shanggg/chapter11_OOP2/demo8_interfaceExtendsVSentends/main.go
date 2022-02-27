package main

import "fmt"

type Monkey struct {
	Name string
}

// 声明接口
type BridAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "生来会爬树..")
}

// LittleMonkey结构体
type LittleMonkey struct {
	Monkey // 继承
}

func (littleM LittleMonkey) Flying() {
	fmt.Println(littleM.Name, "学会了飞翔~")
}

func (littleM LittleMonkey) Swimming() {
	fmt.Println(littleM.Name, "学会了游泳~")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
