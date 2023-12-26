package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归出口
	if root == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val { // p, q 都在 root 的左子树中
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val { // p, q 都在 root 的右子树中
		return lowestCommonAncestor(root.Right, p, q)
	} else { // p, q 分别在 root 的左右子树中
		return root
	}
}

func main() {
	fmt.Println(lowestCommonAncestor(&TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}}, &TreeNode{Val: 5}, &TreeNode{Val: 1}))
}
