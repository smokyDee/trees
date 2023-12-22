package binary

import (
	"testing"
)

func TestInorder(t *testing.T) {
	root := createtree()
	inorder(root)
}

func TestPreorder(t *testing.T) {
	root := createtree()
	preorder(root)
}

func TestPostorder(t *testing.T) {
	root := createtree()
	postorder(root)
}

func createtree() *Node {
	//level 4
	n4 := NewNode(4, nil, nil)
	n6 := NewNode(6, nil, nil)
	n12 := NewNode(12, nil, nil)
	n14 := NewNode(14, nil, nil)

	//level 3
	n1 := NewNode(1, nil, nil)
	n5 := NewNode(5, n4, n6)
	n10 := NewNode(10, nil, nil)
	n13 := NewNode(13, n12, n14)

	//level 2
	n2 := NewNode(2, n1, nil)
	n7 := NewNode(7, n5, nil)
	n11 := NewNode(11, n10, n13)

	//level 1
	n3 := NewNode(3, n2, n7)
	n9 := NewNode(9, nil, n11)

	//root
	n8 := NewNode(8, n3, n9)
	return n8
}
