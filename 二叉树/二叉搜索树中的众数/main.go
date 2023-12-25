package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	res := []int{}

	var pre *TreeNode

	var maxCount, count int

	// 中序遍历
	var vist = func(node *TreeNode) {}
	vist = func(node *TreeNode) {
		if node == nil {
			return
		}

		vist(node.Left)

		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			count = 1
		}

		if count > maxCount {
			maxCount = count
			res = []int{node.Val}
		} else if count == maxCount {
			res = append(res, node.Val)
		}

		pre = node

		vist(node.Right)
	}

	vist(root)
	return res
}

func main() {
	// [1,null,2,2]
	fmt.Println(findMode(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}}))
}
