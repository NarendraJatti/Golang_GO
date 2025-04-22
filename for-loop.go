package main 
import "fmt"

func main() {
	// or i :=1
	for i :=1;i <=5 ;i++ {
		if i ==3 {
			break
			//continue  >>skip current iteration i =3
		}
		fmt.Println(i*i)

		// i+=1

		/*

		i :=1
		for i <=5{
			fmt.Println(i*i)
			i+= 1}
		*/
	}
}