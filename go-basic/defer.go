package main

import "fmt"

func main() {
	// Put in a stack the defer
	// and run in the return of the function
	// This is helpful to close connections and clean envs in the end of the code
	defer fmt.Println("com defer, veio primeiro")
	fmt.Println("Sem defer, veio depois")

}
