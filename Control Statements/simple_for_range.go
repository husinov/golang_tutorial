package main
import "fmt"

func main(){

	players := []string{"Ronaldo", "Messi", "Carvalho"}
	for i, j := range players{
		fmt.Println(i + 1, " - ", j)
	}
}
