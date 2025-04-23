package main

import "fmt"

type Student struct {
	name   string
	rollNo int
}

// Function that accepts a Student struct
func print_in(s Student) {
	fmt.Println("Printing roll number:", s.rollNo)
	fmt.Println("Printing name:", s.name)
}

func main() {
	// Creating a Student instance
	st := Student{
		name:   "naren",
		rollNo: 56,
	}

	// Passing struct to function
	print_in(st)

	// Printing struct directly
	fmt.Printf("Student struct: %+v\n", st)
	fmt.Println("Student name:", st.name)
}
