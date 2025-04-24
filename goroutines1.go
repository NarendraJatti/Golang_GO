package main

import (
	"fmt"
	"net/http"
	"time"
)

func printNumbers() {
    for i := 1; i <= 3; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printNumbers()  // Runs concurrently (Line 1)
    fmt.Println("Main") // Line 2
    time.Sleep(1 * time.Second) // Waits for goroutine (Line 3)
}


func main() {
    for i := 0; i < 100; i++ {
        go http.Get("https://example.com") // Each may block on I/O
    }
    time.Sleep(5 * time.Second)
}