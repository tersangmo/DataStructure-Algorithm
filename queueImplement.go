package main

import "fmt"

func main() {
	var queue []int

	queue = append(queue, 1, 2, 4, 6, 7)

	for len(queue) > 0 {
		// first in first out

		// print the popped element

		out := queue[0]
		fmt.Println(out)

		// update the queue

		queue = queue[1:]
	}
}
