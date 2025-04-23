package main 
import "fmt"

func printName(str string) {
	fmt.Println(str)
}

func printRollNo(rno int) {
	fmt.Println(rno)
}

func printAddress(adr string) {
	fmt.Println(adr)
}

func main() {
	printName("naren")
	defer printRollNo(56)
	printAddress("vijayapura")
	
}