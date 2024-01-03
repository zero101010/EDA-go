package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	contador := 0
	total := 10

	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {
			// Usa o Mutex para trabar uma variável
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			// Usa o Mutex para destravar uma variável
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Valor final", contador)
}
