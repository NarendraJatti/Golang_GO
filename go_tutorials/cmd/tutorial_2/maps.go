package main

import "fmt"

//var myMap2 = map[string]uint8{"naren": 1, "jatti": 2}

func main() {
	var myMap2 = map[string]uint8{"naren": 1, "jatti": 2}

    fmt.Println(myMap2["naren"])
	var  rno,ok = myMap2["john"]

	if ok{
		fmt.Printf("the age is %v",rno)
	}else {
		fmt.Println("invalid name")
	}
      for name:= range myMap2{
		fmt.Printf("name: %v\n",name)
	}
	for key, value := range myMap2 {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
