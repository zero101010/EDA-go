package main

import "fmt"

func square(v *int) {
	*v *= *v
	fmt.Println(*v)
}
func main() {
	a := 4
	square(&a)
	fmt.Println(a)
}
