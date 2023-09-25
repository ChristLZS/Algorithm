package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func visit(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	// 左子树
	res = append(res, visit(root.Left)...)

	// 右子树
	res = append(res, visit(root.Right)...)

	// 根节点
	res = append(res, root.Val)

	return res
}

func postorderTraversal(root *TreeNode) []int {
	return visit(root)
}

func main() {}
