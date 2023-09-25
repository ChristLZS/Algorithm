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

	// 根节点
	res = append(res, root.Val)

	// 右子树
	res = append(res, visit(root.Right)...)

	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	return visit(root)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	inorderTraversal(root)
}
