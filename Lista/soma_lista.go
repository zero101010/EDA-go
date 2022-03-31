package main

import (
	"fmt"
)
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var actualNode *ListNode
	var result *ListNode
	carry := 0
	// Faz a iteração até chegar em um valor null
	for(l1 != nil && l2 != nil) {
		// Faz a soma com o valor da carry
		sum := l1.Val + l2.Val + carry
		// Salva o endereço do novo nó temporário
		tempNode := &ListNode{Val: sum}
		if actualNode == nil {
			// Salva a primeira posição da lista encadeada
			result = tempNode
			actualNode = tempNode
		} else {
			// salva o endereço da memória 
			actualNode.Next = tempNode
			actualNode = actualNode.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	
	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	return result
}

func main() {
	list1 := ListNode{2,&ListNode{2,&ListNode{3,&ListNode{2,nil}}}}
	list2 := ListNode{3,&ListNode{5,&ListNode{1,nil}}}
	addTwoNumbers(&list1, &list2)
}