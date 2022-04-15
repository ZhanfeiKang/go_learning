package main

import "fmt"

func yuejieError() {
	m := map[byte]byte{
		1: 1,
		2: 2,
	}

	a := m[3]
	fmt.Println("a: ", a)

}
func main() {
	yuejieError()
}
