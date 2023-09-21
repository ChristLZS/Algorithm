package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func visit(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// 根节点
	res := make([]int, 0)
	res = append(res, root.Val)

	// 左子树
	if root.Left != nil {
		res = append(res, visit(root.Left)...)
	}

	// 右子树
	if root.Right != nil {
		res = append(res, visit(root.Right)...)
	}

	return res
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, root.Val)
	res = append(res, visit(root.Left)...)
	res = append(res, visit(root.Right)...)
	return res
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(preorderTraversal(root))
}
