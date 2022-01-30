package main

import "fmt"

func hello(one chan bool) {
	fmt.Println("Hello, again!")
	one <- false
}
func main() {
	one := make(chan bool)
	go hello(one)
	<- one
	fmt.Println("main\nFinished!")
}
