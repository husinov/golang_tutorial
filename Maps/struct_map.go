package main

import "fmt"

type student struct {
	age int
	GPA float32
}

func main() {
	Astudent := student{
		20,
		4.13,
	}
	Bstudent := student{
		21,
		4.3,
	}
	Cstudent := student{
		25,
		2.5,
	}
	studentInfo := map[string]student{
		"Nuriddin": Astudent,
		"Carlos":   Bstudent,
		"John":     Cstudent,
	}
	for first_name, info := range studentInfo {
		fmt.Printf("%s is %d years old and he got %.2f GPA.\n", first_name, info.age,
		info.GPA)
	}
}
