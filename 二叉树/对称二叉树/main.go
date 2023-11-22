package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}

	return isSymmetricForTwoNode(root.Left, root.Right)

}

func isSymmetricForTwoNode(p, q *TreeNode) bool {
	// 比较结构
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	// 比较值
	if p.Val != q.Val {
		return false
	}
	return isSymmetricForTwoNode(p.Left, q.Right) && isSymmetricForTwoNode(p.Right, q.Left)
}

func main() {
	// [2,3,3,4,5,5]
	fmt.Println(isSymmetric(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}))

}
