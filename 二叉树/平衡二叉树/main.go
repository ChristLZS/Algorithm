package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// -1 表示不平衡
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := getDepth(root.Left)
	if leftDepth == -1 { // 左子树不平衡
		return -1
	}

	rightDepth := getDepth(root.Right)
	if rightDepth == -1 { // 右子树不平衡
		return -1
	}

	if abs(leftDepth-rightDepth) > 1 { // 左右子树高度差大于1
		return -1
	}

	return max(leftDepth, rightDepth) + 1
}

func isBalanced(root *TreeNode) bool {
	return getDepth(root) != -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(isBalanced(nil))
}
