package tree

import "fmt"

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
func (bt *BinaryTree) PreOrderTraversalByRecursion(root *TreeNode) []int {
	res := make([]int, 0)
	res = preOrderTraversal(root, res)
	return res
}

func preOrderTraversal(root *TreeNode, res []int) []int {
	if root==nil {
		return res
	}
	res = append(res, root.Val)
	res = preOrderTraversal(root.Left, res)
	res = preOrderTraversal(root.Right, res)
	return res
}


// 先序遍历 迭代版
func (bt *BinaryTree) PreOrderTraversalByIteration(root *TreeNode) []int {
	return nil
}

// 中序遍历 递归版
func (bt *BinaryTree) InOrderTraversalByRecursion() []int {
	return nil
}

// 中序遍历 迭代版
func (bt *BinaryTree) InOrderTraversalByIteration() []int {
	return nil
}

// 后序遍历 递归版
func (bt *BinaryTree) PostOrderTraversalByRecursion() []int {
	return nil
}

// 后序遍历 迭代版
func (bt *BinaryTree) PostOrderTraversalByIteration() []int {
	return nil
}