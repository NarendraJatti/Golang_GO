package main

import (
    "fmt"
    "myproject/utils"  // import your local package
)

func main() {
    result := utils.Add(5, 7)
    fmt.Println("Result:", result)
}
