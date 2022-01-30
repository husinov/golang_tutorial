package main

import "fmt"

func main() {
	a := 5
	var b *int
	if b == nil {
		fmt.Println("define b, cause it equals to", b)
	}
	b = &a
	fmt.Printf("b now equals to %v\n", b)
}
