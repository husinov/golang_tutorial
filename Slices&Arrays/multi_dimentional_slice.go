package main
import "fmt"

func main(){
	md_slice := [][]int{{12, 23},
						{43, 56},
						{50, 45}}
						
	fmt.Print("I guess this number is 56:\n", md_slice[1][1])

	if(md_slice[1][1] == 56){
		fmt.Println("\nYes!")
	}else{
		fmt.Println("\nNo!")
	}
}
