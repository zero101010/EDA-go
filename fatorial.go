package main

import "fmt"

func fat(num int) int {
	if num == 1 || num == 0 {
		return 1
	}
	return num * fat(num-1)
}
func main() {

	num := 0

	fmt.Println(fat(num))
}
