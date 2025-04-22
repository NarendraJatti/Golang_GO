  package main 

  import("fmt")

  //below is func signature
  func addNumbers(a int, b int) int { //parms

	sum := a + b
	return sum 
  }

  func main() {
  result := addNumbers(3,7) //fun args
  fmt.Println(result)

  }