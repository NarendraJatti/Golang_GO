package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Print the number of available CPU cores
	fmt.Println("CPU Cores:", runtime.NumCPU())

	// By default, GOMAXPROCS = NumCPU
	runtime.GOMAXPROCS(4) // Manually set (optional)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d running on core\n", id)
		}(i)
	}
	wg.Wait()
}