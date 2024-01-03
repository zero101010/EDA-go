package escrevendo

import "fmt"

const x = 10

// Doc é um monte de coisa nenhuma
func Doc() {
	fmt.Println("Essa função não faz nada e printa Doc")
}

// Mostra a implementação do doc2
func doc2() {
	fmt.Println("Essa função não faz nada e printa doc2")
}

// Demonsta o Doc3 para explicar como faz uma documentação
func Doc3() {
	fmt.Println("Essa função não faz nada e printa Doc3", x)
}
