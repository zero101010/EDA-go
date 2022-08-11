package main

import "fmt"

type Stack struct {
	items []string
}

// Push
func (s *Stack) Push(word string) {
	s.items = append(s.items, word)
}

// // POP
func (s *Stack) Pop() string {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	var arr = []string{}
	word := "viver"
	s := Stack{}
	for _, elem := range word {
		s.Push(string(elem))
	}
	fmt.Println(s.items)
	for i := len(s.items) - 1; i >= 0; i-- {
		unStack := s.Pop()
		arr = append(arr, unStack)
	}
	fmt.Println(arr)

}
