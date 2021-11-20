package main

import "fmt"

// Soma de todos os elementos de um array
// Esse exemplo é para analizar a complexidade do código em Big O
// no nosso caso a função sum_array tem complexidade O(n) quando analizamos o tempo
// por conta do único for

func sum_array(arr [2]int) int {
	sum := 0
	for _, element := range arr {
		sum = sum + element
	}
	return sum
}

func main() {
	numbers := [2]int{1, 2}
	fmt.Println(sum_array(numbers))
}
