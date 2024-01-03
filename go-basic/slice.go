package main

import "fmt"

// Slice is a composite data type, this mean that could be using the primitive types
// Array has defined lenght and slice not
// We can not change the lenght with array but we can change the lenght of slice
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 6}
	fmt.Println(slice)
	slice = append(slice, 5)
	fmt.Println(slice)
	// range
	for i, value := range slice {
		fmt.Println("No indice", i, "Temos o valor:", value)
	}

	// Slice slices
	slices := slice[0:3]
	fmt.Println(slices)
	// Delete slice the 3 index
	// result  should be {1, 2, 3, 6,5}, because we delete index 3 that contains the value 4
	// This 3 poinst(...) that instead of use the slice of ints we will use int by int to add to the other slice
	deleteSlices := append(slice[:3], slice[4:]...)
	fmt.Println(deleteSlices)
}
