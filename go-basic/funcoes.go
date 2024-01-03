package main

import "fmt"

// Parametro variatico
// We can pass one or more values or don't pass any value
func soma(x ...int) (int, int) {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum, len(x)
}

func main() {
	total, tamanho := soma(10, 10, 3)

	fmt.Println(total, tamanho)
}
