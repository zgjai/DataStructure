package tree

import (
	"fmt"
	"testing"
)

func TestEmptyTree(t *testing.T) {
	arr := []int{1, 2, 0, 3}
	bt := newBinaryTree(arr)
	fmt.Println(bt.PostOrderTraversalByIteration())
}

func TestBinaryTree_Balanced(t *testing.T) {
	arr := []int{1,2,3,4,0,5,6,0,0,0,0,0,0,7}
	bt := newBinaryTree(arr)
	fmt.Println(bt.PreOrderTraversalByIteration())
	if !bt.IsBalanced() {
		t.Error("tree is balanced")
	}
}

func TestBinaryTree_NotBalanced(t *testing.T) {
	arr := []int{1,2,3,0,0,5,6,0,0,0,0,7}
	bt := newBinaryTree(arr)
	fmt.Println(bt.PreOrderTraversalByIteration())
	if bt.IsBalanced() {
		t.Error("tree is not balanced")
	}
}
