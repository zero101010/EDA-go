package main

import "fmt"

// return the func argument from another function

func returnFunc(x int) func(int) int {
	return func(i int) int {
		return i * 10
	}
}

func main() {
	x := 5

	y := func(a int) int {
		return a * 10
	}(x)

	fmt.Println(y)

}
