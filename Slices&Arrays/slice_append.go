package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4}
	num = append(num, 5)
	fmt.Println(num)

	var str = []string{"a", "b", "c"}
	str_2 := make([]string, 4, 9)
	str_2 = append(str, "d")
	fmt.Print(str, "\n", str_2)
}
