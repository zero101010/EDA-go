package main

import "fmt"

// 1,2,3,4,5,6,7,8,9,10
// O(log n)
func busca_binaria(arr []int, item int) (int, bool) {
	// inicio do array
	baixo := 0
	// final do array
	alto := len(arr) - 1
	for baixo <= alto {
		// valor do meio é 4
		meio := alto + baixo
		chute := arr[meio]
		if item == chute {
			return arr[meio], true
		}
		if item < chute {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1, false
}
func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(busca_binaria(array, 3))
}
