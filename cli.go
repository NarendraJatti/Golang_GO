// C:\Users\NAREN JATTI\OneDrive\Documents\kodekloud\golang>cli.exe my-giddu
Hello, my-giddu

package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) > 1 {
        fmt.Println("Hello,", os.Args[1])
    } else {
        fmt.Println("Hello, World!")
    }
}