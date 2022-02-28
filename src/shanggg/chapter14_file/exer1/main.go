package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 用于统计结果
type CharCount struct {
	chCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	otherCount int // 记录其他字符的个数
}

func main() {
	// 思路：代开一个文件，创建一个Reader
	// 读取每一行，就去统计该行有多少个英文、数字、空格和其他字符
	// 然后将结果保存到一个结构体
	fileName := "../test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}

	defer file.Close()

	// 定义CharCount实例
	var count CharCount

	reader := bufio.NewReader(file)

	// 开始循环读取文件内容
	fmt.Println()
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(string(str))
		// 遍历 str, 进行统计
		// 为了兼容中文，可以将str转成[]rune
		// str2 := []rune(str)
		for _, v := range []rune(str) {

			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				count.chCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.otherCount++
			}

		}
	}

	// 输出统计的结果看看是否正确
	fmt.Println()
	fmt.Println("字符: ", count.chCount)
	fmt.Println("数字: ", count.NumCount)
	fmt.Println("空格: ", count.SpaceCount)
	fmt.Println("其他: ", count.otherCount)
}
