/*

go tool pprof cpu.prof

(pprof) top
go tool pprof mem.prof

Key pprof Features
Profile Type	Use Case	How to Enable
CPU	Find slow functions	pprof.StartCPUProfile()
Memory (Heap)	Detect memory leaks	pprof.WriteHeapProfile()
Blocking	Find goroutine stalls	runtime.SetBlockProfileRate(1)
Goroutines	Track goroutine count	pprof.Lookup("goroutine")
*/

package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

// computeFibonacci calculates Fibonacci numbers (CPU-intensive)
func computeFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return computeFibonacci(n-1) + computeFibonacci(n-2)
}

// allocateMemory creates memory load
func allocateMemory() {
	var data [][]int
	for i := 0; i < 1000; i++ {
		slice := make([]int, 10000) // Allocate 10k integers
		data = append(data, slice)
		time.Sleep(1 * time.Millisecond) // Simulate work
	}
}

func main() {
	// CPU Profiling
	cpuProfile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer cpuProfile.Close()
	pprof.StartCPUProfile(cpuProfile) // Start CPU profiling
	defer pprof.StopCPUProfile()     // Stop when done

	// Memory Profiling
	memProfile, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer memProfile.Close()

	// Run CPU-heavy task
	go func() {
		log.Println("Fibonacci(40):", computeFibonacci(40))
	}()

	// Run memory-heavy task
	allocateMemory()

	// Write memory profile at the end
	pprof.WriteHeapProfile(memProfile)
}