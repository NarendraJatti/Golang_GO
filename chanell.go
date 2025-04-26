package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		sum :=0
		for i :=0 ;i <100 ;i++ {
			fmt.Println("idx from the first fune",i)
			sum += i
		}
		 c <- sum
	}()

 output := <-c
 fmt.Println("output:",output)
}