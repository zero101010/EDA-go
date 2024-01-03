package main

import "fmt"

// Have the len and the capacity of slice. This help to increase the performance because we know how much of memory will be allocate
func main() {
	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4
	fmt.Println(slice, len(slice), cap(slice))

}
