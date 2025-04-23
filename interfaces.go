package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
}

// Define a struct
type Circle struct {
	radius float64
}

// Circle implements Shape interface by defining Area method
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Another struct
type Rectangle struct {
	width, height float64
}

// Rectangle also implements Shape
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Function that works with any Shape
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 4, height: 3}

	printArea(c) // Works with Circle
	printArea(r) // Works with Rectangle
}
