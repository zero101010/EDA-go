package main

import (
	"encoding/json"
	"fmt"
)

// https://pkg.go.dev/encoding/json
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

// To export struct to json all of keys need to be Uppercase
// For this reason contabancaria
// To transform a json to go struct has this site https://transform.tools/json-to-go
type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	contabancaria float64
}

func main() {
	group := ColorGroup{1, "Igor", []string{"Vermelho", "preto"}}
	fmt.Println(group)
	// Marshal return slice of bytes that's mean if you want to check
	// You can go to ascii and check the code of each letter
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error", err)
	}
	// Byte by Byte
	fmt.Println(b)
	// os.Stdout.Write(b)

	pessoa := Pessoa{"Igor", "Araújo", 26, "Garoto de Programa", 100}

	j, err := json.Marshal(pessoa)
	fmt.Println(j)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(string(j))
}
