package main

import "fmt"

// 声明一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
	Name string
}

// 让Camera实现 Usb 接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (p Camera) Stop() {
	fmt.Println("相机停止工作...")
}

// 计算机
type Computer struct {
}

// 编写一个 Working 方法，接收一个Usb接口类型变量
// 只要实现了Usb接口，所谓实现Usb接口，就是指实现了Usb声明的所有方法
func (c Computer) Working(usb Usb) {

	// 通过usb接口来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {

	// 定义一个Usb接口数组，可以存放phone和camera的结构体变量
	// 这里就体现处多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	fmt.Println(usbArr)
}
