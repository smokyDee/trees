package stack

import (
	"errors"

	"github.com/smokyDee/trees/internal/binary"
)

type StackEntry struct {
	node   *binary.Node
	status string
}

func NewStack() *Stack {
	s := new(Stack)
	s.stacktop = 0
	return s
}

type Stack struct {
	Entry    [20]StackEntry
	stacktop int
}

func (s *Stack) StackEmpty() bool {
	return s.stacktop == 0
}

func (s *Stack) Push(n *binary.Node, status string) {
	s.Entry[s.stacktop].node = n
	s.Entry[s.stacktop].status = status
	s.stacktop++
}

func (s *Stack) Pop() (StackEntry, error) {
	if s.stacktop == 0 {
		return StackEntry{}, errors.New("stack is empty")
	}
	s.stacktop--
	return s.Entry[s.stacktop], nil
}
