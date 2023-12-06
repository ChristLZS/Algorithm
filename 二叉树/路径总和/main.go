package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	curSum := 0
	var vist func(node *TreeNode) bool
	vist = func(node *TreeNode) bool {
		// cur
		if node != nil {
			curSum += node.Val // 计算当前和
		}

		if node.Left == nil && node.Right == nil {
			if curSum == targetSum {
				return true
			}
		}

		// left
		if node.Left != nil {
			if vist(node.Left) {
				return true
			} else {
				curSum -= node.Left.Val // 回溯
			}
		}

		// right
		if node.Right != nil {
			if vist(node.Right) {
				return true
			} else {
				curSum -= node.Right.Val // 回溯
			}
		}

		return false
	}

	return vist(root)
}

func main() {
	// [1,-2,-3,1,3,-2,null,-1]
	// 3
	fmt.Println(hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -1,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -2,
			},
			Right: nil,
		},
	},
		3))
}
