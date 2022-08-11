package main

import "fmt"

type ListNode struct {
	Val  string
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head

	var prev *ListNode
	for {
		if curr == nil {
			return prev
		}
		// passa o valor do próximo para uma variável
		nxt := curr.Next
		// Muda o valor do nó atual que apontava para o próxima para apontar para o anterior
		curr.Next = prev
		// fmt.Println(curr)
		// coloca o valor do atual como anterior
		prev = curr
		// coloca valor do próximo como atual
		curr = nxt

	}
}

func main() {
	list1 := ListNode{"A", &ListNode{"B", &ListNode{"C", nil}}}
	// Output expected [8,9,9,9,0,0,0,1]
	fmt.Println(reverseList(&list1))

}
