package main

import "fmt"

func main(){
	v := 100
	if(v < 50){
		fmt.Println("v is lower than 50")
	}else if(v < 150){
		fmt.Println("v is lower than 150")
	}
	fmt.Printf("v = %d", v)
}
