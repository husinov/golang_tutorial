package main

import (
	"fmt"
)

func call(num *int) {
	*num = 10
}

func main() {
	val := 5
	fmt.Println("value =", val)
	add := &val
	call(add)
	fmt.Println("value now equals to", val)
}
