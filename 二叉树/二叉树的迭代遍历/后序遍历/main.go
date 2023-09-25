package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		// 添加根节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		// 添加左节点
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		// 添加右节点
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}

	// 反转
	reverse(res)

	return res
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[j], nums[i] = nums[i], nums[j]
	}
}

func main() {
	fmt.Println(postorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
