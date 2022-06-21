package main

import (
	"fmt"
)

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {

		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// Exercicio https://www.codewars.com/kata/5663f5305102699bad000056/train/go

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	var arr1 []int
	for _, value := range a1 {
		arr1 = append(arr1, len(value))
	}
	min1, max1 := findMinAndMax(arr1)
	var arr2 []int
	for _, value := range a2 {
		arr2 = append(arr2, len(value))
	}
	min2, max2 := findMinAndMax(arr2)
	fmt.Println("Mínimo do primeiro array:", min1)
	fmt.Println("Máximo do primeiro array:", max1)

	fmt.Println("Mínimo do segundo array:", min2)
	fmt.Println("Máximo do segundo array:", max2)

	result1 := absInt(max1 - min2)
	result2 := absInt(min1 - max2)
	if result1 <= result2 {
		return result2
	}
	if result2 < result1 {
		return result1
	}
	return 1
}
func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
func absInt(x int) int {
	return absDiffInt(x, 0)
}
func main() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	// a1 := []string{}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	fmt.Println(MxDifLg(a1, a2))

}
