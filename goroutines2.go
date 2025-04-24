package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printLine1(){
	defer wg.Done() //Decrements counter by 1 / Will run when function exits
	fmt.Println("Line-1")
	//time.Sleep(10000 * time.Millisecond) //wait-time
	fmt.Println("wait simulation done")
}

func printLine2(){
	// defer.wg.Done() //Missing wg.Done()! // Now called when printLine2 exits
	fmt.Println("Line-2")
}

func main(){
	fmt.Println("main func")
	wg.Add(1) 
	//wg.Add(2) // Set counter to 2
	//wg.Add(2) >gives error
	go printLine1() //Will call Done()
	go printLine2() // Never calls Done()
	wg.Wait() // Wait for all goroutines to finish //// Waits forever (counter stuck at 1)..wait till counter reaches 0


	//time.Sleep(1000 * time.Millisecond) 
}

/*
Execution Flow (Fixed):
wg.Add(2) → Counter = 2.

printLine1() → defer wg.Done() → Counter = 1 when it exits.

printLine2() → defer wg.Done() → Counter = 0 when it exits.

wg.Wait() now unblocks because counter reached 0.
*/