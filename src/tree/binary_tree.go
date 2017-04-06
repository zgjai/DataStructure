package tree

import (
	"fmt"
	"stack"
)

// 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

// 根据输入数组构建一个二叉树，方便测试用，假设0表示这个节点不存在
func initBinaryTreeWithArr(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == 0 {
		return nil
	}
	tn := &TreeNode{Val: arr[0]}
	fmt.Println(tn.Val)
	if len(arr) > 0 {
		tn.Left = initBinaryTreeWithArr(arr[1:])
	}
	if len(arr) > 1 {
		tn.Right = initBinaryTreeWithArr(arr[2:])
	}
	return tn
}

// 先序遍历 递归版
func (bt *BinaryTree) PreOrderTraversalByRecursion() []int {
	res := make([]int, 0)
	res = preOrderTraversal(bt.Root, res)
	return res
}

func preOrderTraversal(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = append(res, node.Val)
	res = preOrderTraversal(node.Left, res)
	res = preOrderTraversal(node.Right, res)
	return res
}

// 先序遍历 迭代版
func (bt *BinaryTree) PreOrderTraversalByIteration() []int {
	if bt.Root == nil {
		return nil
	}
	// 借助stack保存当前节点
	res := []int{}
	s := stack.NewStack(0)
	node := bt.Root
	for node != nil || !s.IsEmpty() {
		if node != nil {
			res = append(res, node.Val)
			s.Push(node)
			node = node.Left
		} else {
			node = s.Pop().(*TreeNode).Right
		}
	}
	return res
}

// 先序遍历 迭代简单版
func (bt *BinaryTree) PreOrderTraversalByIterationEasy() []int {
	if bt.Root == nil {
		return nil
	}
	// 借助stack保存当前节点
	res := []int{}
	s := stack.NewStack(0)
	s.Push(bt.Root)
	// 开始从stack中不断弹出节点，直至stack为空
	for !s.IsEmpty() {
		node := s.Pop().(*TreeNode)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
		res = append(res, node.Val)
	}
	return res
}

// 中序遍历 递归版
func (bt *BinaryTree) InOrderTraversalByRecursion() []int {
	if bt.Root == nil {
		return nil
	}
	res := make([]int, 0)
	res = inOrderTraversal(bt.Root, res)
	return res
}

func inOrderTraversal(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = inOrderTraversal(node.Left, res)
	res = append(res, node.Val)
	res = inOrderTraversal(node.Right, res)
	return res
}

// 中序遍历 迭代版
func (bt *BinaryTree) InOrderTraversalByIteration() []int {
	if bt.Root == nil {
		return nil
	}
	res := make([]int, 0)
	s := stack.NewStack(0)
	node := bt.Root
	for node != nil || !s.IsEmpty() {
		if node != nil {
			s.Push(node)
			node = node.Left
		} else {
			node = s.Pop().(*TreeNode)
			res = append(res, node.Val)
			node = node.Right
		}
	}
	return res
}

// 后序遍历 递归版
func (bt *BinaryTree) PostOrderTraversalByRecursion() []int {
	if bt.Root == nil {
		return nil
	}
	res := make([]int, 0)
	res = postOrderTraversal(bt.Root, res)
	return res
}

func postOrderTraversal(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = postOrderTraversal(node.Left, res)
	res = postOrderTraversal(node.Right, res)
	res = append(res, node.Val)
	return res
}

// 后序遍历 迭代版
func (bt *BinaryTree) PostOrderTraversalByIteration() []int {
	if bt.Root == nil {
		return nil
	}
	res := make([]int, 0)
	s := stack.NewStack(0)
	output := stack.NewStack(0)
	s.Push(bt.Root)
	for !s.IsEmpty() {
		node := s.Pop().(*TreeNode)
		if node.Left != nil {
			s.Push(node.Left)
		}
		if node.Right != nil {
			s.Push(node.Right)
		}
		output.Push(node)
	}
	for !output.IsEmpty() {
		res = append(res, output.Pop().(*TreeNode).Val)
	}
	return res
}
