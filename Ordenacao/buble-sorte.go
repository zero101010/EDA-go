package main

import (
	"fmt"
)

func main() {
	var arr = []int{25, 100, 0, 25, 100}
	fmt.Println(arr)
	fmt.Println(bubleSort(arr))
}

// O(n)
func bubleSort(array []int) []int {
	fmt.Println(len(array))
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; i++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
