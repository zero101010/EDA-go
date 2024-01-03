//https://leetcode.com/problems/binary-search/?envType=study-plan&id=algorithm-i

package main

import "fmt"

func binarySearch(nums []int, target int) int {
	index := 0
	total_index := len(nums) - 1
	index_value := nums[index]

	for index_value != target {
		middle := (len(nums)) / 2
		fmt.Println("Middle", middle)
		prev := nums[:middle]
		prox := nums[middle:]
		if target > prox[0] {
			nums = prox
			fmt.Println("Maior que o target")
			fmt.Println(nums)
		}
		if target < prox[0] {
			nums = prev
			fmt.Println("Menor que o target")
			fmt.Println(nums)
		}
		if target == prox[0] {
			return
		}
	}
	return -1
}

func main() {
	var array = []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(binarySearch(array, 3))
}

//  0 1 2 | 3 4 5
// -1 0 3 | 5 9 12
// indice 3

// -1 | 0 3
// indice 2

// 0 | 3
// indice 1
