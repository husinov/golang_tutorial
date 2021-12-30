package main

import "fmt"

func main(){

	v := 20
	if(v > 5){
		if(v > 10){
			fmt.Println("v is greater than 10")
		}
	}
	fmt.Printf("v equals to %d", v)
}
