package main

import "fmt"

func main() {
	// Use functions as expressions. Search some use cases to see the best whey to use this
	x := 10
	y := func(x int) int {
		return x * 10
	}
	fmt.Println(y(x))
}
