package main

import "fmt"

func main() {
    chars := []string{"a", "b", "c", "d"}
    charChan := make(chan string, len(chars)) // Buffered channel with capacity for all items

    for _, char := range chars {
        select {
        case charChan <- char:
            fmt.Println("Sent:", char)
        }
    }

    // Now receive the values
    close(charChan) // Close the channel when done sending
    for char := range charChan {
        fmt.Println("Received:", char)
    }
}

/* 

package main 

import (
    "fmt"
    "time"
)

func main() {
    chars := []string{"a", "b", "c", "d"}
    mychan := make(chan string)

    // Start a receiver goroutine
    go func() {
        for char := range mychan {
            fmt.Println("Received:", char)
        }
    }()

    // Send values
    for _, char := range chars {
        select {
        case mychan <- char:
            fmt.Println("Sent:", char)
        }
    }

    // Give some time for the receiver to process
    time.Sleep(100 * time.Millisecond)
    close(mychan)
}

*/