package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// 每次处理一层
	for len(queue) > 0 {
		tmpRes := make([]int, 0)
		tmpQueue := make([]*TreeNode, 0)

		// 处理一层
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			tmpRes = append(tmpRes, cur.Val)
			if cur.Left != nil {
				tmpQueue = append(tmpQueue, cur.Left)
			}
			if cur.Right != nil {
				tmpQueue = append(tmpQueue, cur.Right)
			}
		}
		res = append(res, tmpRes)
		queue = tmpQueue
	}

	return res
}

func main() {
	fmt.Println(levelOrder(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
