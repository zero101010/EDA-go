package main

import (
	"strings"
)

func reverseString(s string) string {
	var arr = []string{}
	// For each
	for i := len(s) - 1; i >= 0; i-- {
		arr = append(arr, string(s[i]))
	}

	return strings.Join(arr, "")
}

func reverseString2(s string) (result string) {

	for _, elem := range s {
		result = string(elem) + result
	}
	return result
}
func main() {
	print(reverseString2("teste"), "\n")

}
