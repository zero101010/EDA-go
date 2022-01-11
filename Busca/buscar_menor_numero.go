package main

import "fmt"

// O(n) pois passa por todos os itens para fazer a verificacao
func BuscaMenor(arr []int) int {
	// salva o menor valor
	menor := arr[0]
	for indice, element := range arr {
		if arr[indice] < menor {
			menor = element
		}
	}
	return menor
}
func main() {

	var array = []int{5, 6, 7, 8, 2, 9, 10, 7}
	fmt.Println(buscaMenor(array))
}
