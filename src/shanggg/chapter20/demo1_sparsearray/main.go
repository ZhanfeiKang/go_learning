package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}

func writeArr(sparseArr []ValNode) {
	filePath := "./chessmap.data"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// 输出稀疏数组
	fmt.Println("正在将稀疏数组写入文件......")
	for i, valNode := range sparseArr {
		str := fmt.Sprintf("%d %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
		writer.WriteString(str)
	}

	writer.Flush()
}

func readArr() (chessMap2 [11][11]int) {
	fmt.Println("正在读取数组文件......")
	filePath := "./chessmap.data"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var (
		str1   = ""
		strArr []string
		row    = 0
		col    = 0
		val    = 0
	)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		str1 = strings.TrimRight(str, " \r\n")
		strArr = strings.Split(str1, " ")
		row, err = strconv.Atoi(strArr[1])
		if err != nil {
			fmt.Println("strconv.Atoi err: ", err)
		}
		col, err = strconv.Atoi(strArr[2])
		if err != nil {
			fmt.Println("strconv.Atoi err: ", err)
		}
		val, err = strconv.Atoi(strArr[3])
		if err != nil {
			fmt.Println("strconv.Atoi err: ", err)
		}
		if strArr[0] != "0" {
			chessMap2[row][col] = val
		}
		// fmt.Println(strArr)
	}

	return
}

func main() {

	// 1.先创建一个原始的数组
	var chessMap [11][11]int

	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 白子

	// 2.输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组。
	// 思路
	// 1) 遍历 chessMap，如果我们发现有一个元素的值不为0
	// 2) 创建一个node结构体，将其放入到对应的切片中即可

	var sparseArr []ValNode

	// 标准的一个稀疏数组应该还有一个表示记录原始的二维数组的规模（行和列，默认值）
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个 ValNode 值结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 将这个稀疏数组，存盘 ./chessmap.data
	writeArr(sparseArr)

	// 恢复原始数组
	chessMap2 := readArr()
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
