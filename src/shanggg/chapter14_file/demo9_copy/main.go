package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 自己编写一个函数，接受两个文件路径 srcFileName  dstFileName
func CopyFile(dstFileName string, srcFilName string) (written int64, err error) {
	srcfile, err := os.Open(srcFilName)
	if err != nil {
		fmt.Println("open fil err: ", err)
		return
	}
	defer srcfile.Close()

	// 通过srcfile，获取到Reader
	reader := bufio.NewReader(srcfile)

	// 打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer dstFile.Close()

	// 通过dstFile,获取到Writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)

}

func main() {
	// 将 ./dir1/daisy.jpg 拷贝到 ./dir2/zhuzhu.jpg

	// 调用CopyFile完成文件拷贝
	srcFile := "./dir1/daisy.jpg"
	dstFile := "./dir2/zhuzhu.jpg"

	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("Copy complete~")
	} else {
		fmt.Println("copy err: ", err)
	}
}
