package main

import "fmt"

type Usb interface {
	Say()
}

type Stu struct {
}

func (stu *Stu) Say() {
	fmt.Println("Say()")
}

func main() {
	var stu Stu = Stu{}
	// var u Usb = stu // 错误！会报Stu类型没有实现Usb接口
	var u Usb = &stu
	u.Say()
	fmt.Println("here", u)
}
