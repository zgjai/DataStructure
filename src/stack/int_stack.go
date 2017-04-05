package stack

import (
	"fmt"
)

type IntStack struct {
	*Stack
}

func NewIntStack(c int) *IntStack {
	return &IntStack{NewStack(c)}
}

func (s *IntStack) Push(ele int) {
	s.Stack.Push(ele)
}

func (s *IntStack) Pop() int {
	return s.Stack.Pop().(int)
}

func (s *IntStack) Top() int {
	return s.Stack.Top().(int)
}

func (s *IntStack) IsEmpty() bool {
	return s.Stack.IsEmpty()
}

func (s *IntStack) Depth() int {
	return s.Stack.Depth()
}

func (s *IntStack) String() string {
	return fmt.Sprintf("%d", s.items)
}
