package main

import "fmt"

func main() {
	x := 10
	// Get the value of x address in memory and use as y value
	y := &x
	// dereferencing is basically get the value of memory adress
	z := *y
	fmt.Println(z)
	fmt.Println(*&x)

	recebeValor(x)
	fmt.Println(x)
	recebePonteiro(&x)
	fmt.Println(x)

}

func recebeValor(x int) {
	x++
	fmt.Println("Valor dentro da função", x)
}
func recebePonteiro(x *int) {
	*x++
}
