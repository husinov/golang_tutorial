package main

import "fmt"

func main() {
	a := make(chan int)
	if a != nil {
		fmt.Println("a is defined")
	}
}
