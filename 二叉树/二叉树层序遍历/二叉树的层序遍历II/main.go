package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				tmpQueue = append(tmpQueue, cur.Left)
			}
			if cur.Right != nil {
				tmpQueue = append(tmpQueue, cur.Right)
			}
			tmpRes = append(tmpRes, cur.Val)
		}

		res = append(res, tmpRes)
		queue = tmpQueue
	}

	// 反转
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func main() {
	fmt.Println(levelOrderBottom(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
