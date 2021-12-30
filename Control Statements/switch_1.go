package main
import "fmt"

func main(){

	switch player := 1; player{
	case 1:
		fmt.Println("Ronaldo")
	case 2:
		fmt.Println("Messi")
	case 3:
		fmt.Println("Lewandowski")
	}
}
