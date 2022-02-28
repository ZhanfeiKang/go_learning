package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 拷贝文件将 ../test.txt 文件内容内容导入到 ./kk.txt

	// 1.首先将 ../test.txt 内容读取到内存
	// 2.将对读取到的内容写入 ./kk.txt

	file1Path := "../test.txt"
	file2Path := "./kk.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Println("write file err: ", err)
		return
	}
}
