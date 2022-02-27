package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	// 概念：file的三种叫法
	// 1.file对象
	// 2.file指针
	// 3.file文件句柄
	file, err := os.Open("../test.txt")
	if err != nil {
		fmt.Println("open file err: ", err)
	}

	// 输出下文件，看看文件是什么,看出file就是一个指针 *File
	fmt.Println("file=%v", file)

	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err: ", err)
	}
}
