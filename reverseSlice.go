package main

import "fmt"

// without temporary slice
func reverse1(s []int) []int {
	for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}
	return s
}

//with temporary slice

func reverse2(s []int) []int {
	s1 := make([]int, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		s1 = append(s1, s[i])
	}
	return s1
}

func main() {
	slice1 := []int{5, 3, 1, 8, 9}

	fmt.Println("The reverse slice without using temporary slice is:", reverse1(slice1))
	slice2 := []int{2, 1, 7, 8, 9}
	fmt.Println("The reverse slice with using temporary slice is:", reverse2(slice2))

}
