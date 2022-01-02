package main
import "fmt"

func add(first, second int)(int){

	result := first + second
	return result
}
func main(){
	fmt.Printf("Result: %d", add(10, 15))
}
