package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func main() {
	fmt.Println(invertTree(nil))
}
