func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1) // Add before goroutine
        go func(i int) {
            defer wg.Done() // Done at the end (deferred)
            // Do work
            fmt.Println(i)
        }(i)
    }
    
    wg.Wait() // Wait for all goroutines to complete
    fmt.Println("All done")
}