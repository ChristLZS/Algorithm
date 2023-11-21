package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root
	stack = append(stack, cur)
	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil { // 添加节点
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}

			// 标记节点
			stack = append(stack, cur)
			stack = append(stack, nil)

			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}

		} else { // 处理节点
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
		}
	}

	return res
}

func main() {
	fmt.Println(inorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
