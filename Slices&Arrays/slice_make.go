package main
import "fmt"

func main(){
	var slice_num = make([]int, 3, 6)

	fmt.Println(slice_num, len(slice_num), cap(slice_num))
}
