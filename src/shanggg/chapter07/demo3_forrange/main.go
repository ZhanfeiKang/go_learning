package main

import "fmt"

func main() {
	var heros [3]string = [3]string{"宋江", "吴用", "卢俊义"}

	for index, hero := range heros {
		fmt.Printf("heros[%v]:%v\n", index, hero)
	}

	for _, hero := range heros {
		fmt.Printf("%v\n", hero)
	}
}
