package main

import "fmt"


func sliceToChannel(nums []int) <- chan int {
	out :=make(chan int)

	go func() {
		for _,n := range nums {
			out <- n
		}
		close(out)
	}()

	return out

}

func main() {

	//inPut
	nums := []int{2, 3, 4, 5}
	//stage-1
	dataChannel := sliceToChannel(nums)

	// stage-2
	finalChannel := sq(dataChannel)

	// stage-3
	for n := range finalChannel{
		fmt.Println(n)
	}


	// fmt.Println(nums)
}