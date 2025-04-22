package main

import "fmt"

func main() {

	var languages [2]string = [2]string{"Kannada", "English"}
	fmt.Println(languages)
	fmt.Println(languages[0])
	for i := 0; i < len(languages); i++ {
		fmt.Println(languages[i])
	}

	marks := [3]int{10, 20, 30}
	fmt.Println(marks)

	pets := [...]string{"dogs", "cats"}
	fmt.Println(pets)
}
