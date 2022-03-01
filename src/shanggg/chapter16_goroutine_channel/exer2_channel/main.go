package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// var perChan chan Person
	perChan := make(chan Person, 11)

	var person Person

	fmt.Println("---------------------------------------")
	for i := 0; i < 10; i++ {

		person.Name = "person" + strconv.Itoa(i)
		person.Age = rand.Intn(100)
		person.Address = "address" + strconv.Itoa(i)
		// fmt.Println(person)
		perChan <- person
	}

	fmt.Println(len(perChan))
	close(perChan)
	fmt.Println("---------------------------------------")
	for v := range perChan {
		fmt.Println(v)
	}
	fmt.Println("---------------------------------------")
}
