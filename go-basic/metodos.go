package main

import "fmt"

type pessoa struct {
	name string
	age  int
}

// func (receiver) identifier(parameters) (returns){code}
// receiver ---> It will say who it's the owner of the method
func (p pessoa) oibomdia() {
	fmt.Println("Olá", p.name)
}

func main() {
	pessoa1 := pessoa{"Igor", 26}

	pessoa1.oibomdia()
}
