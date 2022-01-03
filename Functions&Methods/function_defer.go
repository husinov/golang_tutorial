package main
import "fmt"

func multiply(a, b int){

	fmt.Println(a * b)
}

func main(){

	defer multiply(10, 15)
	fmt.Println("Multiply")	
}
