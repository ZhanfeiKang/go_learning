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

func (p Phone) Call() {
	fmt.Println("手机在打电话...")
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

func (c Computer) Working(usb Usb) {

	usb.Start()
	// 如果usb是指向Phone结构体变量，则还需要调用Call方法
	// 类型断言..【注意体会】
	if phone, ok := usb.(Phone); ok {
		fmt.Println(usb.(Phone))
		phone.Call()
	}
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

	// 遍历usbArr
	// Phone还有一个特有的方法call()，请遍历Usb数组，如果是Phone变量，
	// 除了调用Usb接口声明的方法外，还需要调用Phone特有方法 call.  => 类型断言
	var computer Computer
	for _, v := range usbArr {
		fmt.Println("----------------")
		computer.Working(v)
	}

}
