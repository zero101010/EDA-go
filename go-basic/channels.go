package main

import "fmt"

func main() {
	channel := make(chan int)
	go func() {

		channel <- 42
		channel <- 43

	}()
	// Channel work as a queue to async, including with concorrencia
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
