package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewIntStack(0)
	a := [...]int{2, 3, 4, 5, 7, 1, 2, 10, 20, 38, 79}
	for _, i := range a {
		s.Push(i)
	}
	if s.Depth() != len(a) {
		t.Error("push fault")
	}
}

func TestPop(t *testing.T) {
	s := NewIntStack(0)
	a := [...]int{2, 3, 4, 5, 7, 1, 2, 10, 20, 38, 79}
	for _, i := range a {
		s.Push(i)
	}
	if s.Pop() != 79 {
		t.Error("pop fault")
	}
	if s.Top() != 38 {
		t.Error("pop fault")
	}
	if s.Pop() != 38 {
		t.Error("pop fault")
	}
}

func TestDepth(t *testing.T) {
	s := NewIntStack(10)
	a := [...]int{2, 3, -4, 5, 7, -1, 2, 10, 20, 38, 79}
	for _, i := range a {
		s.Push(i)
	}
	if s.Depth() != len(a) {
		t.Error("depth fault")
	}
}

func TestArrayStack(t *testing.T) {
	s := NewIntStack(5)
	s.Push(-10)
	if s.Top() != -10 {
		t.Error("top fault")
	}

	if s.IsEmpty() {
		t.Error("should be false")
	}
}
