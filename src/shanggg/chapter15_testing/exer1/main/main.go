package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `hson:"age"`
	Skill string `json:"skill"`
}

func (monster *Monster) Store() bool {

	fmt.Println("monster: ", *monster)
	str, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("marshal err: ", err)
		return false
	}
	filePath := "monster.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(string(str) + "\n")
	writer.Flush()
	return true
}

func (monster *Monster) Restore() bool {
	filePath := "./monster.txt"
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return false
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	str := ""
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str += line
	}
	// fmt.Println(str)
	err = json.Unmarshal([]byte(str), monster)
	if err != nil {
		fmt.Println("unmarshal err", err)
		return false
	}

	fmt.Println("monster: ", *monster)
	return true
}

func main() {
	var monster Monster = Monster{
		Name:  "牛魔王",
		Age:   500,
		Skill: "牛头拳",
	}
	monster.Store()
	monster.Restore()
}
