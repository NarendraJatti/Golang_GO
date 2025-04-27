package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// ch <- 5

	// fmt.Println(<- ch)
	wg.Add(2)
	//belwo channel is revding the value
	// recevie ONLY chanel
	go func(ch <-chan int,wg * sync.WaitGroup){
		val,isChanelOpen := <-myCh
		
		fmt.Println(isChanelOpen)
		fmt.Println(val)


		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh,wg)

	//go func(ch chan int,wg * sync.WaitGroup)
	//below sending teh value the value
	//send ONLY chanbel below
	go func(ch chan<- int,wg * sync.WaitGroup){
		myCh <-0			
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		// close(myCh)
		wg.Done()
	}(myCh,wg)


	wg.Wait()
}