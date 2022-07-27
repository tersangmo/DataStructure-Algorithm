package main

import "fmt"

func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		} else if len(right) == 0 {
			return append(result, left...)
		} else if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func MergeSort(value []int) []int {
	if len(value) <= 1 {
		return value
	}
	mid := len(value) / 2
	left := MergeSort(value[:mid])
	right := MergeSort(value[mid:])
	return Merge(left, right)
}

func main() {
	value := []int{9, 4, 3, 6, 1, 2, 1, 0, 5, 7, 8}

	fmt.Printf("Original value is :%v\n"+"New value is:%v\n", value, MergeSort(value))
}
