package main

import "fmt"

func main() {
	a := i()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

// We create a closure, this mean that only will increase the value inside of the function
// For this reason when i pass "b" and "a" have different values.
// They increase but only in each closure

func i() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
