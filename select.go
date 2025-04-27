package main

import "fmt"

func main() {

	firstchannel := make(chan string)
	secondchannel := make(chan string)

	go func() {
		firstchannel <- "datapoint-1"
	}()

	go func() {
		secondchannel <- "datapoint-2"
	}()

	select {
	case msgFrom1stchannel := <-firstchannel:
		fmt.Printf(msgFrom1stchannel)
	case msgFrom2stchannel := <-secondchannel:
		fmt.Println(msgFrom2stchannel)

	}
}