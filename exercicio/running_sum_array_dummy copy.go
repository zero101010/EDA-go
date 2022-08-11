package main

import "fmt"

// https://leetcode.com/problems/running-sum-of-1d-array/

func sumArray(arr []int) int {
	sum := 0
	for _, element := range arr {
		sum = sum + element
	}
	return sum
}

func runningSum(nums []int) []int {
	array := []int{}
	for i, _ := range nums {
		res := sumArray(nums[:i+1])
		array = append(array, res)
	}
	return array
}

func main() {
	arr := []int{1, 2, 3}
	// fmt.Println(arr[:3])
	fmt.Println(runningSum(arr))
}
