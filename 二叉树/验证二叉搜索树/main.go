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
// 如果是二叉搜索树，那么中序遍历的结果应该是一个升序数组
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var pre = math.MinInt64
	var vist func(root *TreeNode)
	vist = func(root *TreeNode) {

		if root.Left != nil {
			vist(root.Left)
		}

		if root.Val > pre {
			pre = root.Val
		} else {
			pre = math.MaxInt64
		}

		if root.Right != nil {
			vist(root.Right)
		}
	}

	vist(root)

	return pre != math.MaxInt64
}

func main() {
	// [5,4,6,null,null,3,7]
	fmt.Println(isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}
