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

func runTest(numCores int) time.Duration {
	// Set number of cores to use
	runtime.GOMAXPROCS(numCores)
	fmt.Printf("\n=== Running with %d core(s) ===\n", numCores)
	fmt.Printf("Actual CPU cores available: %d\n", runtime.NumCPU())
	
	ch := make(chan string)
	start := time.Now()
	
	// Start workers
	go dowork(1, time.Second*2, ch)
	go dowork(2, time.Second*1, ch)
	go dowork(3, time.Second*3, ch)
	
	// Main work
	for i := 0; i < 3; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("Main working at %s\n", time.Now().Format("15:04:05.000"))
	}
	
	// Wait for completion
	for i := 0; i < 3; i++ {
		<-ch
	}
	
	duration := time.Since(start)
	fmt.Printf("Total time with %d core(s): %v\n", numCores, duration)
	return duration
}

func main() {
	// Test with different core counts
	duration1 := runTest(1)  // Single core (concurrency only)
	duration2 := runTest(2)  // Two cores (some parallelism)
	duration4 := runTest(4)  // More cores (more potential parallelism)
	
	fmt.Printf("\nPerformance comparison:\n")
	fmt.Printf("Single core: %v\n", duration1)
	fmt.Printf("Two cores:   %v (%+.2f%%)\n", duration2, 
		float64(duration1-duration2)/float64(duration1)*100)
	fmt.Printf("Four cores:  %v (%+.2f%%)\n", duration4,
		float64(duration1-duration4)/float64(duration1)*100)
}