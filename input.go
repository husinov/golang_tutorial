package main
import "fmt"

func main(){
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hi %s. Welcome to my GitHub account.", name)
}
