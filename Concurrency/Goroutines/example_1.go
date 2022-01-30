package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(time.Millisecond)
}

func hello() {
	fmt.Println("Hello")
}
