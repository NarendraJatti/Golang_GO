package main

import (
	"errors"
	"fmt"
)

func main() {

	printavalue := "hi go,naren"
	printMe(printavalue)

	var numerator int =20
	var denominator = 0
	var result_final,remainder,err = intDivison(numerator,denominator)
	//fmt.Println("result",result_final)
	if err!=nil {
		fmt.Println(err.Error())
	}else if  remainder == 0{
		fmt.Printf("the result is %v",result_final)
	} else {
	fmt.Printf("the reuslt is %v and remainder is %v",result_final,remainder) }
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