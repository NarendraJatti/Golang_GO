package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// ch <- 5

	// fmt.Println(<- ch)
	wg.Add(2)
	go func(ch chan int,wg * sync.WaitGroup){

		fmt.Println(<-myCh)
		wg.Done()
	}(myCh,wg)

	go func(ch chan int,wg * sync.WaitGroup){

		myCh <- 5

		wg.Done()
	}(myCh,wg)


	wg.Wait()
}