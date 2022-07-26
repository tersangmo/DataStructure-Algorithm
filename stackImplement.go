package main

import "fmt"

func main() {
	var stack []int

	stack = append(stack, 4, 5, 7, 8)

	for len(stack) > 0 {
		//print the top
		top := len(stack) - 1
		fmt.Println(stack[top])

		//pop the top
		stack = stack[:top]
	}
}
