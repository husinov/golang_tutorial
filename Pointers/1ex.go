package main

import "fmt"

func main() {
	a := 10
	var b *int = &a
	fmt.Printf("Type of b is %T\n", b)
	fmt.Println(b)
}
