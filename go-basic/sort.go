package main

import (
	"fmt"
	"sort"
)

//https://pkg.go.dev/sort
// Time complexity is O(n*log(n)) is worst than linear
func main() {
	ss := []string{"f", "h", "s", "a", "b"}
	sort.Strings(ss)
	fmt.Println(ss)
	i := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Ints(i)
	fmt.Println(i)
}
