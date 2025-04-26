package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
func compute(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	start := time.Now()
	// Simulate CPU-intensive work
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Printf("Goroutine %d finished after %v\n", id, time.Since(start))
}
func main() {
    numCPU := runtime.NumCPU()
    fmt.Printf("Number of CPUs: %d\n", numCPU)
    runtime.GOMAXPROCS(numCPU)
// Parallel version
    start := time.Now()
    var wg sync.WaitGroup
    for i := 0; i < 8; i++ {
        wg.Add(1)
        go compute(&wg, i) // <- The 'go' keyword makes it parallel
    }
    wg.Wait()
    fmt.Printf("Parallel total time: %v\n", time.Since(start))
// Sequential version
    start = time.Now()
    for i := 0; i < 8; i++ {
        var wgSeq sync.WaitGroup
        wgSeq.Add(1)
        compute(&wgSeq, i)  // <- No 'go' keyword
        wgSeq.Wait()
    }
    fmt.Printf("Sequential total time: %v\n", time.Since(start))}




// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// func compute(wg *sync.WaitGroup, id int) {
// 	defer wg.Done()
	
// 	start := time.Now()
// 	// Simulate CPU-intensive work
// 	sum := 0
// 	for i := 0; i < 100000000; i++ {
// 		sum += i
// 	}
// 	fmt.Printf("Goroutine %d finished after %v\n", id, time.Since(start))
// }

// func main() {
// 	// Print number of available CPUs
// 	numCPU := runtime.NumCPU()
// 	fmt.Printf("Number of CPUs: %d\n", numCPU)
	
// 	// Allow Go to use all available CPUs
// 	runtime.GOMAXPROCS(numCPU)
	
// 	var wg sync.WaitGroup
	
// 	for i := 0; i < 8; i++ {
// 		wg.Add(1)
// 		go compute(&wg, i)
// 	}
	
// 	wg.Wait()
// }

// func compute(wg *sync.WaitGroup, id int) {
//     defer wg.Done()
//     start := time.Now()
//     sum := 0
//     for i := 0; i < 100000000; i++ {
//         sum += i
//     }
//     // Get OS thread ID
//     threadID := runtime.LockOSThread()
//     runtime.UnlockOSThread()
//     fmt.Printf("Goroutine %d (Thread %d) finished after %v\n", id, threadID, time.Since(start))
// }