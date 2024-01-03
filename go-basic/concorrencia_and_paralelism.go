package main

import (
	"fmt"
	"sync"
	"time"
)

// O primeiro CPU dual core veio em 2006
// O Go foi a primeira lingaguem criada com mult-cores em mente
// Concorrencia é um padrão de desing que várias coisas sejam
// executadas de formas independentes
// Concorrencia seria fazer um conjunto de coisas em seguência mas não ao mesmo tempo, como conversar, digitar, voltar a conversar e depois digitar
// Paralelismo quando o código executa várias coisas ao mesmo tempo entre os cores, exemplo
// Conversar e teclar ao mesmo tempo
// Por padrão o golang usa o máximo de paralelismo possível

// Se eu rodar sem o WaitGroup a func main vai terminar antes da func main
// Por isso usamos o WaitGroup
var wg sync.WaitGroup

func main() {
	// O 1 é pq estamos usando somente 1 go routine
	wg.Add(2)
	go func1()
	go func2()
	// Manda a func main esperar até a goroutine ser executada
	wg.Wait()
}

func func1() {
	for i := 0; i < 20; i++ {
		fmt.Println("Func1:", i)
		time.Sleep(20)
	}
	// Manda o sinal de que a func da go routine foi executada
	wg.Done()
}
func func2() {
	for i := 0; i < 20; i++ {
		fmt.Println("Func2:", i)
		time.Sleep(20)
	}
	wg.Done()
}
