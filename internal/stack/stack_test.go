package stack

import (
	"fmt"
	"testing"

	"github.com/smokyDee/trees/internal/binary"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(binary.NewNode(2, nil, nil), "1")
	stack.Push(binary.NewNode(3, nil, nil), "2")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
