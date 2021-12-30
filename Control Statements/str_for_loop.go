package main
import "fmt"

func main(){

	for i, j := range "Hello"{
		fmt.Printf("The index number of %U is %d\n", j, i)
	}
}
