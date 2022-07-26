package main

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a, b := 1, 2

	fmt.Println("Minimum number is :", Min(a, b))
	fmt.Println("Maximum number is :", Max(a, b))

}
