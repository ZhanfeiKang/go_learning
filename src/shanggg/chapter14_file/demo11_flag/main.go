package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义几个变量，勇于接受命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	// &user: 就是接收用户命令行中输入的 -u 后面的参数值
	// "u": 就是 -u 指定参数
	// " ": 默认值
	// "用户名，默认为空": 说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认为3306")

	// 这里有一个非常重要的动作
	// 从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行
	flag.Parse()

	// 输出结果
	fmt.Println("user: ", user)
	fmt.Println("pwd: ", pwd)
	fmt.Println("host: ", host)
	fmt.Println("port: ", port)
}
