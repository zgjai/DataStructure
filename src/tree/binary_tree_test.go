package tree

import (
	"fmt"
	"testing"
)

func TestEmptyTree(t *testing.T) {
	arr := []int{1, 2, 0, 3}
	bt := BinaryTree{}
	bt.Root = initBinaryTreeWithArr(arr)
	fmt.Println(bt.Root.Val)
}
