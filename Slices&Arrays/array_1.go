package main

import "fmt"

func main(){
	var arr[3]string

	arr[0] = "First"
	arr[1] = "Second"
	arr[2] = "Third"

	for key, _ := range arr{
		fmt.Print(key, "       ")
	}
	fmt.Println()
	fmt.Println(arr)
}
