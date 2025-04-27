package main

import (
	"fmt"
	"time"
)

func dowork(d time.Duration) {
	fmt.Println("doing work")
	time.Sleep(d)
	fmt.Println("work is done")
}

func main(){
	start := time.Now()
	dowork(time.Second *2)
	fmt.Printf("work took %v seconds\n",time.Since(start))
}