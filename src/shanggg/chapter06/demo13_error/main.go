package main

import (
	"errors"
	"fmt"
)

func test() {
	// 使用 defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到异常
			fmt.Println("err=", err)
			// 这里就可以将错误信息发送给管理员。。。
			fmt.Println("发送邮件给admin@kkite.com")
		}
	}()
	num1 := 10
	num2 := 0

	res := num1 / num2
	fmt.Println("res=", res)
}

// 函数去读取一个配置文件的init.conf的信息
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config2.ini" {
		// 读取..
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config.ini")
	if err != nil {
		// 如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行...")
}

func main() {
	test()

	fmt.Println("main()下面的代码...1")

	test02()
	fmt.Println("main()下面的代码...2")

}
