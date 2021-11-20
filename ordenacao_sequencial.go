package main

import "fmt"

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
	var array = []int{10, 2, 4, 5, 1, 6}
	fmt.Println(ordenacaoSequencial(array))
}
