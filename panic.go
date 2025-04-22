package main

import "fmt"

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Result:", a/b) // panic if b is 0
}

func main() {
	safeDivide(10, 0)
	fmt.Println("Program continues...")
}
