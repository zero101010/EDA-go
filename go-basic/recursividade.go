package main

import "fmt"

func fatorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * fatorial(x-1)
}
func main() {
	fmt.Println(fatorial(2))
}
