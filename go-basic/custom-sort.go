package main

import (
	"fmt"
	"sort"
)

type Carro struct {
	nome     string
	potencia int
	consumo  int
}

type OrdenacaoPorPotencia []Carro

func (x OrdenacaoPorPotencia) Len() int           { return len(x) }
func (x OrdenacaoPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (a OrdenacaoPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func main() {
	carros :=
		[]Carro{
			{"Chevete", 500, 5},
			{"Porche", 300, 5},
			{"Fusca", 5, 30},
		}
	fmt.Println(carros)
	sort.Sort(OrdenacaoPorPotencia(carros))
	fmt.Println(carros)

}
