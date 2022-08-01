package main

import "fmt"

/*
this function is to remove element while
maintaining the order of the slice
*/
func Remove1(s []int, i int) []int {
	// copy returns the number of elements copied which is from
	//i+1 to the last position of element
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

/*
this function is to remove element
without maintaining the order of the slice
*/
func Remove2(s []int, i int) []int {
	// it involves moving the last element into the gap
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	data := []int{5, 6, 7, 8, 9}

	fmt.Println(Remove1(data, 2))
	fmt.Println(Remove2(data, 2))

}
