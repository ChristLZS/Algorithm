package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	var visit func(node *TreeNode) *TreeNode
	visit = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val < low {
			return visit(node.Right)
		} else if node.Val > high {
			return visit(node.Left)
		}

		node.Left = visit(node.Left)
		node.Right = visit(node.Right)

		return node
	}

	return visit(root)
}

func main() {
	fmt.Println(trimBST(nil, 0, 0))
}
