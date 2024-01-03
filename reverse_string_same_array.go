// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func main() {

	arr := []int{5, 5, 2, 6, 2}
	init := 0
	end := len(arr) - 1
	temValueInit := 0
	temValueEnd := 0
	for {
		if init == end {
			fmt.Println(arr)
			return
		}
		temValueInit = arr[init]
		temValueEnd = arr[end]
		arr[init] = temValueEnd
		arr[end] = temValueInit
		init = init + 1
		end = end - 1
		fmt.Println("init", init)
		fmt.Println("end", end)
	}
}
