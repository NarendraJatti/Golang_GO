package main

import "fmt"

// Higher-order function: takes another function as a parameter
func applyFunc(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

// Simple function to pass
func add(a int, b int) int {
	return a + b
}

func main() {
	result := applyFunc(10, 5, add)
	fmt.Println("Result:", result) // Output: Result: 15
}
