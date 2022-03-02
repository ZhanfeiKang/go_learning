package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// net.Dial("tcp", "127.0.0.1:8888")
	conn, err := net.Dial("tcp", "192.168.56.1:8888")
	if err != nil {
		fmt.Println("client dial err: ", err)
		return
	}
	// fmt.Println("conn successful, conn: ", conn)

	defer conn.Close()

	// 功能1: 客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) // os.Stain 代表标准输入[终端]

	var sumByte int
	for {
		// 从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err: ", err)
		}

		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			break
		}
		// 再将line 发送给服务器
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err: ", err)
		}
		sumByte += n
	}

	fmt.Printf("客户端发送了 %d 字节数据，并退出...", sumByte)

}
