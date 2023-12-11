package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var min = math.MaxInt64
	var pre *TreeNode
	var vist func(root *TreeNode)
	vist = func(root *TreeNode) {
		if root.Left != nil {
			vist(root.Left)
		}

		if pre != nil && root.Val-pre.Val < min {
			min = root.Val - pre.Val
		}
		pre = root

		if root.Right != nil {
			vist(root.Right)
		}
	}

	vist(root)

	return min
}

func main() {
	// [4,2,6,1,3]
	fmt.Println(getMinimumDifference(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}))
}
