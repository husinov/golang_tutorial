package main

import "fmt"

type Player struct {
	name    string
	goals   int
	assists int
}

func (p Player) info() {
	fmt.Printf("%s scored %d goals and assisted %d.\n", p.name, p.goals, p.assists)
}

func main() {
	aPlayer := Player{
		"Ronaldo",
		30,
		10,
	}
	bPlayer := Player{
		"Modric",
		10,
		30,
	}
	aPlayer.info()
}
