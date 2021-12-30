package main

import "fmt"

func main(){

	var name interface{}
	switch type_name := name.(type){
	case bool:
		fmt.Println("Boolean")
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float")
	default:
		fmt.Printf("The type of name is %T", type_name)
	}
}
