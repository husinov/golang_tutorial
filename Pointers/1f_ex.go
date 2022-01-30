package main

import "fmt"

func main() {
	b := hi()
	fmt.Println("b =", *b)
}

func hi() *int {
	a := 5
	return &a
}
