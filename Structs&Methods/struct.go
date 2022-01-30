package main
import "fmt"

type Student struct{
	full_name string
	age int
	region string
	GPA float64

}

func main(){
	var aStudent = Student{"Nuriddin Husinov", 20, "Navoiy", 4.2}
	fmt.Println(aStudent)
}
