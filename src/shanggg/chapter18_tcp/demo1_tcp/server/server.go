package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	// 这里我们循环的接收客户端发送的数据
	defer conn.Close() // 关闭con

	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// conn.Read(buf)
		// 1.等待客户端通过conn发送信息
		// 2.如果客户端没有write[发送], 那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端: %s 发送信息: \n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn读取
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出...")
				return
			}

			fmt.Println("服务器的Read err: ", err)
			return
		}

		// 3.显示客户端发送的内容到服务器的终端
		// buf[:n] 注意：这里一定要带n，读到n，因为底层会在后面加内容
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	// net.Listen("tcp", "127.0.0.1:8888")
	// 1.tcp: 表示使用网络协议是tcp
	// 2.127.0.0.1:8888：表示在本地监听 8888 端口
	// listen, err := net.Listen("tcp", "127.0.0.1:8888")
	listen, err := net.Listen("tcp", "192.168.56.1:8888")
	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}
	defer listen.Close() // 延时关闭listen

	// 循环等待客户端来连接
	for {
		// 等待客户端来连接
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err: ", err)
		} else {
			// fmt.Printf("Accept() successful, conn: %v, 客户端ip: %v\n", conn, conn.RemoteAddr().String())
			fmt.Printf("Accept() successful, 客户端ip: %v\n", conn.RemoteAddr().String())
		}
		// 这里准备起一个协程，为客户端服务
		go process(conn)
	}

	// fmt.Println("listen successful: ", listen)
}
