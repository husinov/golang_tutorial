package main
import "fmt"

func main(){

	new_map := map[int]string{
		1 : "Musk",
		2 : "Gates",
		3 : "Arno",
	}
	for key, value := range new_map{
		fmt.Println(key, "-", value)
	}
}
