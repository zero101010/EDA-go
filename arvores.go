package main

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

func (n *Node) Insert(value int) {
	if n.Value < value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {

		}

	}
}

func main() {

	root := &Node{Value: 10}
	root.Insert(12)
}
