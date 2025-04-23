package main 

import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v",&i ,&i)
	fmt.Printf("%T %v \n",*(&i),*(&i))
	s := "hello"
	var b *string = &s 
	fmt.Println(b)
	var a = &s 
	fmt.Println(a)

	c := &s 
	fmt.Println(c)
}