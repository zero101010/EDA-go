package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O(n long n)
func QuickSort(arr []int) []int {

	if len(arr) < 2 {
		// fmt.Println("len:", len(arr))
		// fmt.Println("Conteúdo do Array:", arr)
		return arr
	} else {
		// Pego primeiro elemento como pivo para comparar
		pivo := arr[0]
		// fmt.Println("Pivo:", pivo)
		var maiores []int
		var menores []int
		for i := 1; i <= (len(arr) - 1); i++ {
			if pivo <= arr[i] {
				maiores = append(maiores, arr[i])
			}
			if arr[i] <= pivo {
				menores = append(menores, arr[i])
			}
		}
		// fmt.Println("Menores:", menores)
		// fmt.Println("Maiores:", maiores)
		var result []int
		result = append(result, QuickSort(menores)...)
		result = append(result, pivo)
		result = append(result, QuickSort(maiores)...)
		return result
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	arr := rand.Perm(1000000)
	// fmt.Println(arr)
	fmt.Println(QuickSort(arr))
}
