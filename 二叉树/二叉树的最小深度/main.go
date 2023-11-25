package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	res := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		res++
		tmpQueue := []*TreeNode{}
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return res
			}

			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}

			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}

		queue = tmpQueue
	}

	return res
}

func main() {
	fmt.Println(minDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))
}
