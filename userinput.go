package main 

import "fmt"

func main() {
	// var name string
	// var is_muggle bool
	// fmt.Print("enter your name:")
	// fmt.Scanf("%s",&name)
	// fmt.Scanf("%s %t",&name,&is_muggle)
	//fmt.Println("hey there,",name)
	// fmt.Println(name,is_muggle)

	var a string 
	var b int 
	fmt.Print("enter a string and a number")
	count,err := fmt.Scanf("%s %d",&a, &b)
	fmt.Println("count: ",count)
	fmt.Println("error: ",err)
	fmt.Println("a:",a)
	fmt.Println("b:",b)
}