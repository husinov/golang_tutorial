package main
import ("fmt"
		"strings"
)

func joinstr(element ...string)(string){
	return strings.Join(element, " - ")
}
func main(){
	fmt.Println(joinstr("Nuriddin", "Husinov"))
}
