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
	count := 1
	carry := 0
	// Faz a iteração até chegar em um valor null das duas listas 
	for(l1 != nil && l2 != nil) {
		// Faz a soma com o valor da carry
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		if sum >= 10 {
			sum = sum % 10
		}
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

	// Faz a iteração até chegar em um valor null da lista 1
	for (l1 != nil){
		sum := l1.Val + carry
		carry = sum / 10
		if sum >= 10 {
			sum = sum % 10
		}
		count = count+1
		tempNode := &ListNode{Val: sum}
		actualNode.Next = tempNode
		actualNode = actualNode.Next
		l1 = l1.Next
	}
		// Faz a iteração até chegar em um valor null da lista 2
	for (l2 != nil){
		sum := l2.Val + carry
		carry = sum / 10
		if sum >= 10 {
			sum = sum % 10
		}
		tempNode := &ListNode{Val: sum}
		actualNode.Next = tempNode
		actualNode = actualNode.Next
		l2 = l2.Next
	}
	if carry > 0 {
		tempNode := &ListNode{Val: carry}
		actualNode.Next = tempNode
	}
	return result
}

func main() {
	list1 := ListNode{9,&ListNode{9,&ListNode{9,&ListNode{9,&ListNode{9,&ListNode{9,&ListNode{9,nil}}}}}}}
	list2 := ListNode{9,&ListNode{9,&ListNode{9,&ListNode{9,nil}}}}
	// Output expected [8,9,9,9,0,0,0,1]
	firstNode := addTwoNumbers(&list1, &list2)
	for firstNode!=nil {
		fmt.Println(firstNode.Val)
		firstNode = firstNode.Next
	}
}