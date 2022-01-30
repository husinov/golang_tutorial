package main

import "fmt"

func main() {
	a := 7
	b := &a
	fmt.Println("Address of a is", b)
	fmt.Println("a equals to", *b)
}
