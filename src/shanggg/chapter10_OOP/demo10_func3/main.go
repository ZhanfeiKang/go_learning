package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c *Circle) area2() float64 {
	// (*c).Radius = 10
	c.Radius = 10
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var c Circle = Circle{5}
	area := c.area()
	// area2 := (&c).area2()
	area2 := c.area2()
	fmt.Println("area1 :", area)
	fmt.Println("area2 :", area2)
	fmt.Println("c.Radius :", c.Radius)
}
