package main
import "fmt"

type Player struct{
	name, team, country  string
	height_cm, weight_kg int
}
func main(){
	var aPlayer Player = Player{"Ronaldo", "Man Utd", "Portugal", 187, 84}
	var bPlayer = Player{"Messi", "PSG", "Argentina", 171, 68}
	cPlayer := Player{"Lewandovski", "Bayern Munich", "Poland", 185, 80}

	fmt.Printf("Height of %s is %d.\n", aPlayer.name, aPlayer.height_cm)
	fmt.Printf("%s was bor in %s.\n", bPlayer.name, bPlayer.country)
	fmt.Printf("%s plays for %s and his weight is %d kg.\n", cPlayer.name, cPlayer.team, cPlayer.weight_kg)
}
