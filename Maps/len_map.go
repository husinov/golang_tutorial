package main

import "fmt"

func main() {
	teams := map[string]int{
		"Real Madrid": 13,
		"Barcelona":   5,
		"Milan":       8,
	}
	fmt.Println("Length of map equals:", len(teams))
}
