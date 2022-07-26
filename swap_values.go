package main

import "fmt"

func swap(a, b int) []int {
	a, b = b, a
	return []int{a, b}
}

func main() {
	// tuple assignment, short declared
	a, b := 5, 8

	//function call to swap
	result := swap(a, b)

	fmt.Println(result)
}
