package main

import (
	"fmt"
	"sort"
)

// This solution only worked if has only one number repeated
func repeat(arr []int) int {
	for i, elem := range arr {
		if elem == arr[i+1] {
			return elem
		}
	}
	return -1
}
func main() {
	var arr = []int{6, 4, 3, 2, 6, 8, 1}
	sort.Ints(arr)
	fmt.Println(repeat(arr))

}
