package main

import (
	"errors"
	"fmt"
)

func main() {

	printavalue := "hi go,naren"
	printMe(printavalue)

	var numerator int =20
	var denominator = 5
	var result_final,remainder,err = intDivison(numerator,denominator)
	//fmt.Println("result",result_final)
	//switch by default implements breaks
	switch{
		case err!=nil:
			fmt.Println(err.Error())
		case  remainder == 0:
			fmt.Printf("the result is %v",result_final)
		default:
		fmt.Printf("the reuslt is %v and remainder is %v",result_final,remainder)  }
	switch remainder{
	case 0:
		fmt.Printf("the divsion was exact")

	}
}
func printMe(printavalue string) {
	fmt.Println(printavalue)

}
func intDivison(numerator int ,denominator int) (int,int,error) {

	var err error 
	if denominator == 0{
		err = errors.New("can not divide by zero")
		return 0,0,err
	}
	var result int = numerator/denominator
	var remainder int = numerator % denominator
return result ,remainder,err 
}