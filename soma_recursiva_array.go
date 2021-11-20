package main

import "fmt"

func RemoveIndex(arr []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}

func sumArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	length := len(arr)
	value := arr[length-1]
	newArray := RemoveIndex(arr, length-1)
	return value + sumArray(newArray)
}
func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumArray(array))
}
