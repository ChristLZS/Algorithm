package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		sum := 0.0
		count := 0
		tmpQueue := make([]*TreeNode, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			sum += float64(cur.Val)
			count++
			if cur.Left != nil {
				tmpQueue = append(tmpQueue, cur.Left)
			}
			if cur.Right != nil {
				tmpQueue = append(tmpQueue, cur.Right)
			}
		}

		sum /= float64(count)
		res = append(res, sum)
		queue = tmpQueue
	}

	return res
}

func main() {
	fmt.Println(averageOfLevels(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
