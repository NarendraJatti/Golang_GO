// main.go
package main // Required for executables

import "fmt"

// Optional initialization function
func init() {
    fmt.Println("Initializing...")
}

// Required entry point
func main() {
    fmt.Println("Hello, World!")
}