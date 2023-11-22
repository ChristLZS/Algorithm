package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		max := queue[0].Val
		for _, v := range queue {
			if v.Val > max {
				max = v.Val
			}
			if v.Left != nil {
				tmpQueue = append(tmpQueue, v.Left)
			}
			if v.Right != nil {
				tmpQueue = append(tmpQueue, v.Right)
			}
		}

		queue = tmpQueue
		res = append(res, max)
	}

	return res
}

func main() {
	fmt.Println(largestValues(nil))
}
