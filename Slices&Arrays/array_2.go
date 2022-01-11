package main
import "fmt"

func main(){
	num := [4]uint8{1, 2, 3, 4}

	for key, value := range num{
		fmt.Printf("Number in index %d is %d\n", key, value)
	}
}
