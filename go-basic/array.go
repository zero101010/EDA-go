// https://www.youtube.com/watch?v=i_3O4ooSVCM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=59
package main

import "fmt"

var x [5]int

func main() {
	// for default the value not defined in the array will has the value 0
	// [1 10 0 0 0]
	// An array has only the defined lenght
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0])
	fmt.Println(x)
	fmt.Println(len(x))

}
