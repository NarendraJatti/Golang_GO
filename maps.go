 package main 
 import("fmt")

 func main() {

	codes := map[string]string{"en":"english","fr":"french"}
	fmt.Println(codes)
	fmt.Println(codes["en"])
	value,found := codes["en"]
	fmt.Println(value ,found)
	value, found = codes["kn"] //observ the syntax
	fmt.Println(value ,found)

	for key,value := range codes {
		fmt.Println(key,"=>",value)
	}
 }