package main 

import "fmt"
import "strconv"

func main() {
	var i int = 90
	var f float64 = float64(i)

	var s string = strconv.Itoa(i) //convert int to string
	fmt.Printf("%.2f\n",f)
	fmt.Printf("%q",s)
}