package main

import "fmt"

func send(s chan<- int) {
	s <- 42
}

func receive(s <-chan int) {
	fmt.Println("O valor recebido foi", <-s)
}

func myloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}
func main() {
	// Os Canais podem ser bidirecionais ou somente enviar e receber
	// A setinha sempre indica o que o canal vai fazer
	// send channel <- significa que o channel vai colocar a info do canal
	// receive <- channel quer dizer que vamos retirar o valor do chanel
	channel := make(chan int)
	// O valor de Channel é compartilhado por 2 funções concorrentes
	go send(channel)
	receive(channel)
	// Usando range e close para trabalhar com os canais
	c := make(chan int)
	go myloop(10, c)
	for v := range c {
		fmt.Println("Recebi do canal", v)
	}
}
