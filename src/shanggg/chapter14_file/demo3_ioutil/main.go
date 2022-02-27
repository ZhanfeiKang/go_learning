package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 使用ioutil.ReadFile一次性将文件读取到位
	file := "./test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err: ", err)
	}

	// 八度渠道的内容显式到终端
	// fmt.Println(content) // []byte
	fmt.Println(string(content)) // []byte

	// 我们没有显式的Open文件，因此也不需要显式的Close文件
	// 因为，文件的Open和Close被封装到 ReadFile 函数内部
}
