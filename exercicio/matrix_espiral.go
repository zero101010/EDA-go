package main

import (
	"fmt"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
func generateMatrix(n int) [][]int {
	square := n * n
	top := 0
	bottom := n
	left := 0
	right := n
	array := makeRange(1, square)
	result := make([][]int, n)
	// Tratar o problema de setar com o make somente a parte externa do array
	for i := range result {
		result[i] = make([]int, n)
	}
	for top < bottom && left < right {
		// left to right
		for i := left; i < right; i++ {
			result[top][i] = array[0]
			array = remove(array, 0)
		}
		top = top + 1
		// top to bottom
		for i := top; i < bottom; i++ {
			result[i][right-1] = array[0]
			array = remove(array, 0)
		}
		right = right - 1
		// right to left
		for i := right; left < i; i-- {
			result[right][i-1] = array[0]
			array = remove(array, 0)
		}
		bottom = bottom - 1
		// the problem be here
		// 2 1 é o que precisamos no últim bot to top

		// bottom to top
		for i := bottom; i > top; i-- {
			fmt.Println("top:", i)
			fmt.Println("bottom:", bottom)
			fmt.Println("right:", right)
			fmt.Println("left:", left)
			result[i][left] = array[0]
			fmt.Println("linha:", i)
			fmt.Println("Coluna:", left)
			fmt.Println("value:", result[i][right-bottom])
			fmt.Println("--------------------------------")

			array = remove(array, 0)
		}
		left = left + 1

	}
	for i := range result {
		fmt.Println(result[i])
	}
	// fmt.Println("Expiral matrix:",result)

	return result
}

func main() {
	generateMatrix(5)
	// [[1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]]
}
