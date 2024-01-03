package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	contador := 0
	total := 100

	wg.Add(100)

	for i := 0; i < total; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Valor final", contador)
}
