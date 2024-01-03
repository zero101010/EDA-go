package main

import "fmt"

type pessoa struct {
	name string
	age  int
}

type profissional struct {
	// is not mandatory set name value in this case like pessoa pessoa
	// this call anonymous structs
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{name: "Igor", age: 26}
	prof := profissional{
		pessoa:  pessoa1,
		titulo:  "DevOps",
		salario: 15000,
	}
	fmt.Println(pessoa1)
	fmt.Println(prof)
}
