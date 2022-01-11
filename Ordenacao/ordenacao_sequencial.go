package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O(n) pois passa por todos os itens para fazer a verificacao
func BuscaMenor(arr []int) (int, int) {
	// salva o menor valor
	menorValue := arr[0]
	var index int
	for indice, element := range arr {
		if arr[indice] < menorValue {
			menorValue = element
			index = indice
		}
	}
	return menorValue, index
}

func RemoveIndex(arr []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}

// O(nˆ2)
func ordenacaoSequencial(arr []int) []int {
	var ordernedArray []int
	var minValue int
	var indexToRemove int
	fmt.Println(arr)
	for range arr {
		minValue, indexToRemove = BuscaMenor(arr)
		ordernedArray = append(ordernedArray, minValue)
		arr = RemoveIndex(arr, indexToRemove)
		fmt.Println(arr)
		fmt.Println("Remove value:", minValue)
	}
	return ordernedArray
}
func main() {
	rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	arr := rand.Perm(1000000)
	fmt.Println(ordenacaoSequencial(arr))
}
