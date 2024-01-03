package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/valid-anagram/
func IsAnagram(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	w1 := map[byte]int{}
	w2 := map[byte]int{}

	for i, _ := range word1 {
		w1[word1[i]]++
	}
	fmt.Print(w1)
	for i, _ := range word2 {
		w2[word2[i]]++
	}
	isEqual := reflect.DeepEqual(w1, w2)

	return isEqual
}

// 13:42
func main() {
	a := "saudade"
	b := "aduesad"
	print(IsAnagram(a, b))
}
