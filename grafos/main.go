package main

import (
	"container/list"
	"fmt"
)

type Pessoa struct {
	amigos []string
}

func Vendedor(name string) bool {
	if name[len(name)-1:] == "m" {
		return true
	} else {
		return false
	}
}

type Queue []interface{}

func (self *Queue) Push(x interface{}) {
	*self = append(*self, x)
}

func (self *Queue) Pop() interface{} {
	h := *self
	var el interface{}
	l := len(h)
	el, *self = h[0], h[1:l]
	// Or use this instead for a Stack
	// el, *self = h[l-1], h[0:l-1]
	return el
}

func NewQueue() *Queue {
	return &Queue{}
}

func main() {
	// Tabela hash
	eu := Pessoa{[]string{"alice", "bob", "claire"}}
	bob := Pessoa{[]string{"anuj", "peggy"}}
	alice := Pessoa{[]string{"peggy"}}
	claire := Pessoa{[]string{"thom", "jonny"}}
	// // Fila
	// q := NewQueue()
	queue := list.New()
	queue.PushBack(eu)
	queue.PushBack(bob)
	queue.PushBack(alice)
	queue.PushBack(claire)
	// fmt.Println(queue.Front().Value)
	for queue.Len() > 0 {
		e := queue.Front()
		pessoa :=
			fmt.Println(e.Value)
		queue.Remove(e)
	}

}
