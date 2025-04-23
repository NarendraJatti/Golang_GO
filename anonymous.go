package main 

import ("fmt")


func main() {

	x := func(a int, b int) int {
		return a * b 
	} (20,30)
	fmt.Println("T \n",x)
	fmt.Println(x)
}