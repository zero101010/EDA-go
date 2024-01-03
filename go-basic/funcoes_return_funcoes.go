package main

import "fmt"

// We can return a func of func. This mean that we can return other function in one function, and not only the primitive type
// Check the example bellow and see how we can instance the function
func returnFunc() func(int) int {

	return func(i int) int {
		return i * 100
	}
}

func main() {
	x := returnFunc()
	y := (x(3))
	fmt.Println(y)
	fmt.Println(returnFunc()(6))

}
