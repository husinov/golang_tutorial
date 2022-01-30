package main

import "fmt"

func main() {
	slice_words := []string{"Hello.", "How", "are", "you?"}

	for i := 0; i < len(slice_words); i++ {
		fmt.Println(slice_words[i])
	}
	fmt.Println(slice_words)
}
