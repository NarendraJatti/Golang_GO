// Non-Optimized Code (Wastes Memory)
How to Check for Heap Allocations
Use go build -gcflags="-m" to see escape analysis:
/*
Why It’s Inefficient:
Dynamic Resizing Overhead:

numbers := []int{} starts with zero capacity.

Each append() may trigger a new memory allocation + copy when capacity is exceeded.

This leads to O(N²) time complexity in the worst case.

Unnecessary Heap Allocations:

Since the slice grows unpredictably, it may end up on the heap (slower than stack).

*/

package main

import "fmt"

func getNumbers() []int {
    numbers := []int{}  // Slice created with no initial capacity
    for i := 0; i < 1000; i++ {
        numbers = append(numbers, i)  // Reallocates memory frequently
    }
    return numbers
}

func main() {
    nums := getNumbers()
    fmt.Println("First 5:", nums[:5])
}

//Memory-Optimized Version

package main

import "fmt"

func getNumbers() []int {
    numbers := make([]int, 0, 1000)  // Pre-allocates memory (len=0, cap=1000)
    for i := 0; i < 1000; i++ {
        numbers = append(numbers, i)  // No reallocations needed
    }
    return numbers
}

func main() {
    nums := getNumbers()
    fmt.Println("First 5:", nums[:5])
}


/*

Why It’s Optimized:
Pre-Allocation (make([]int, 0, 1000))

Sets initial capacity = 1000, so append() never reallocates.

Runs in O(N) time (faster).

Fewer GC (Garbage Collector) Pressures:

No intermediate arrays are discarded, reducing GC workload.

Better Cache Locality:

Contiguous memory block improves CPU cache efficiency.

When to Optimize?
Use pre-allocation (make) when you know the expected size of slices/maps.

Avoid unnecessary heap allocations (e.g., pointers in hot loops).

Profile with pprof if performance matters.

*/