package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil // 文件或者目录存在
	}
	if os.IsNotExist(err) {
		return false, nil // 文件不存在
	}
	return false, err // 其他错误
}

func main() {
	// filepath1 := "../test.txt"
	filepath1 := "../tessdasdat.txt"
	flag, err := PathExists(filepath1)
	if flag {
		fmt.Println("文件存在")
	} else {
		fmt.Println("文件不存在，或存在其他错误. err: ", err)
	}

}
