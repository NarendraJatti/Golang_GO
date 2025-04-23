package main

import "fmt"

// Define a struct
type Person struct {
	name string
	age  int
}

// Define a method on the Person struct
func (p Person) greet() {
	fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old.")
}

func main() {
	// Create a Person
	p := Person{name: "Narendra", age: 24}

	// Call the method
	p.greet()
}
