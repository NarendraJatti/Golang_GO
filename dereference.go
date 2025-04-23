package main 
import "fmt"

func main() {
	s := "hello"
	fmt.Printf("%T %v \n",s,s)
	ps := &s 
	*ps = "world"
	//fmt.Printf(s)
	fmt.Printf("%T %v \n",s,s)
}