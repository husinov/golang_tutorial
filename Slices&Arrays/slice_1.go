package main
import "fmt"

func main(){
	players := [7]string{"Ronaldo", "Messi", "Carvajal", "Alaba", "Navas",
	"Ter Stegen", "Ramos"}

	slice_players := players[0:2]

	fmt.Println(slice_players)
	fmt.Printf("Length of this slice equals to %d\n", len(slice_players))
	fmt.Printf("Capacity of slice is %d\n", cap(slice_players))
}
