package main

import (
	"fmt"
	"runtime"
	"time"
)

func dowork(id int, d time.Duration, ch chan string) {
	start := time.Now()
	fmt.Printf("Worker %d starting at %s\n", id, start.Format("15:04:05.000"))
	time.Sleep(d)
	end := time.Now()
	fmt.Printf("Worker %d finished at %s (after %v)\n", id, end.Format("15:04:05.000"), d)
	ch <- fmt.Sprintf("Worker %d completed", id)
}

func main() {
	// Print number of available CPU cores
	fmt.Printf("Available CPU cores: %d\n", runtime.NumCPU())
	
	// Create a channel
	ch := make(chan string)
	
	// Start timestamp
	mainStart := time.Now()
	fmt.Printf("Main started at %s\n", mainStart.Format("15:04:05.000"))
	
	// Launch 3 workers concurrently
	go dowork(1, time.Second*2, ch)
	go dowork(2, time.Second*1, ch)
	go dowork(3, time.Second*3, ch)
	
	// Main does other work
	for i := 0; i < 5; i++ {
		fmt.Printf("Main doing other work #%d at %s\n", 
			i+1, time.Now().Format("15:04:05.000"))
		time.Sleep(500 * time.Millisecond)
	}
	
	// Wait for all workers to complete
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Printf("Main received: %s at %s\n", 
			msg, time.Now().Format("15:04:05.000"))
	}
	
	// Total duration
	fmt.Printf("Total execution time: %v\n", time.Since(mainStart))
}