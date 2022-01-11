package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	fmt.Println(array)
	// Aplico recursivamente o array inferior
	first := MergeSort(array[:len(array)/2])
	// Aplico recursivamente o array Superior
	second := MergeSort(array[len(array)/2:])

	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	return Merge(first, second)
	// return array
}

// O(n long n)
func Merge(first []int, second []int) []int {
	var final []int
	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			fmt.Println(first[i], "é menor que", second[j])
			final = append(final, first[i])
			i++
		} else {
			fmt.Println(first[i], "é maior que", second[j])
			final = append(final, second[j])
			j++
		}
	}
	// Caso um dos arrays termine e saia da condição do for anterior ele realoca o que sobrou na ordem do arrray, pois foram ordenados anteriormente
	for ; i < len(first); i++ {
		fmt.Println("Sobrou", first[i])
		final = append(final, first[i])
	}
	for ; j < len(second); j++ {
		fmt.Println("Sobrou", second[j])
		final = append(final, second[j])
	}
	return final
}
func main() {
	rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	arr := rand.Perm(3)
	fmt.Println(MergeSort(arr))
}

// referencia https://qvault.io/golang/merge-sort-golang/
