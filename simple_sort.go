package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	arr := rand.Perm(10)
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)
}
