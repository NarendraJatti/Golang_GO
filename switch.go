 package main 

 import "fmt"

 func main() {
	var i int = 800
	// var a,b int = 10,2o
	switch i {
		case 800:
			fmt.Println("i is 800")
			fallthrough //force this block to print
		case 100,200: 
			fmt.Println("is is eithe 100 or 200")
			
		default:
			fmt.Println("i is neither 0,100 or 200")
	}
 }