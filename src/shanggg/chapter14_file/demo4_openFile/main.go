package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1.打开文件
	filePath := "./abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}

	// 及时关闭file句柄，防止内存泄漏
	defer file.Close()

	// 2.写入
	str := "hello kkite.\r\n"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存的，因此在调用writerString方法时，
	// 其实内容是先写入到缓存的, 所以需要调用Flush方法，将缓存的数据真正写入到文件中，
	// 否则文件中会没有数据！！！
	writer.Flush()
}
