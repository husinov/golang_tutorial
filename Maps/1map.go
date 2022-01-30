package main

import "fmt"

func main() {
	players := map[string]int{
		"Ronaldo": 37,
		"Messi":   35,
		"Benzema": 33,
	}
	for key, _ := range players {
		fmt.Println(key)
	}
	fmt.Printf("Benzema is %d old\n", players["Benzema"])

	fmt.Println(players)
}
