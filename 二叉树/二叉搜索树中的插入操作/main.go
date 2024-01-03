package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if nil == root {
		return &TreeNode{Val: val}
	}

	var vist = func(node *TreeNode) {}
	vist = func(node *TreeNode) {
		if val > node.Val {
			if node.Right == nil {
				// 插入右子树
				node.Right = &TreeNode{Val: val}
			} else {
				// 遍历右子树
				vist(node.Right)
			}
		} else {
			if node.Left == nil {
				// 插入左子树
				node.Left = &TreeNode{Val: val}
			} else {
				// 遍历左子树s
				vist(node.Left)
			}
		}
	}

	vist(root)

	return root
}

func main() {
	fmt.Println(insertIntoBST(nil, 0))
}
