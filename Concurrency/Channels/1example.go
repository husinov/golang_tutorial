package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("Please, define your channel!")
	}
	a = make(chan int)
	fmt.Printf("Type of a is %T\n", a)
}
