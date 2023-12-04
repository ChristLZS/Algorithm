package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	// vist 用于递归遍历二叉树
	var vist func(root *TreeNode, path string)
	vist = func(root *TreeNode, path string) {
		if root.Left == nil && root.Right == nil {
			res = append(res, path+fmt.Sprintf("%d", root.Val))
			return
		}

		if root.Left != nil {
			vist(root.Left, path+fmt.Sprintf("%d->", root.Val))
		}
		if root.Right != nil {
			vist(root.Right, path+fmt.Sprintf("%d->", root.Val))
		}
	}

	vist(root, "")

	return res
}

func main() {
	fmt.Println(binaryTreePaths(nil))
}
