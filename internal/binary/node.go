package binary

import (
	"fmt"
)

func NewNode(value int, left *Node, right *Node) *Node {
	return &Node{
		value: value,
		left:  left,
		right: right,
	}
}

type Node struct {
	left  *Node
	right *Node
	value int
}

func preorder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	preorder(n.left)
	preorder(n.right)
}

func inorder(n *Node) {
	if n == nil {
		return
	}
	inorder(n.left)
	fmt.Println(n.value)
	inorder(n.right)
}

func postorder(n *Node) {
	if n == nil {
		return
	}
	postorder(n.left)
	postorder(n.right)
	fmt.Println(n.value)
}
