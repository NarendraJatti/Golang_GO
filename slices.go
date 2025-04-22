package main
import "fmt"

func main() {

	array := [6]int{1,7,10,1,35,23}
	slice_1 := array[0:2]
	fmt.Println(slice_1)
}