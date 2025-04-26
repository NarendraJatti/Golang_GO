package main

import (
	"fmt"
	"time"
)

func main() {
	go Line_1("Line1")
	Line_1("Line2")

	//fmt.Scanln()


}

func Line_1(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Microsecond)
		fmt.Println(s)
	}
}