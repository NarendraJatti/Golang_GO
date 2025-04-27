package main

import (
	"fmt"
	"time"
)

func dowork(d time.Duration, ch chan string) {
	fmt.Println("doing work")
	time.Sleep(d)
	fmt.Println("work is done")
	ch <- "work completed"  // Send completion message
}

func main() {
	ch := make(chan string)
	
	go dowork(time.Second*2, ch)  // Run concurrently
	go dowork(time.Second*4, ch)

	//dowork(time.Second*2)  // Blocks main() for 2 seconds
	
	// Do other work in main while dowork runs
	fmt.Println("Main can do other things now")
	
	// Wait for completion message
	msg := <-ch
	fmt.Println("Received:", msg)
}