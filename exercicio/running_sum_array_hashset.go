package main

import "fmt"

// https://leetcode.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	array := []int{}
	m := make(map[int]int)
	for i, _ := range nums {
		if i == 0 {
			m[i] = nums[i]
			array = append(array, m[i])
		} else {
			m[i] = m[i-1] + nums[i]
			array = append(array, m[i])
		}
	}
	return array
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(runningSum(arr))
}
