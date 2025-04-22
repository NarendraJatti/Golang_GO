package main 
import (
    "fmt"
    "reflect"
)


func main() {

	var grades int = 42 
	var message string = "hello word"

	// can not use Println for belwo,use Printf
	fmt.Printf("variable grades = %v is of type %T \n",grades,grades)
	fmt.Printf("variable message = %v is of type %v \n",message, reflect.TypeOf(message))

	fmt.Printf("type: %v \n",reflect.TypeOf(11.05))
}