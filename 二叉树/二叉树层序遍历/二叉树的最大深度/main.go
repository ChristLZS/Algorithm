package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	res := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left != nil {
				tmpQueue = append(tmpQueue, v.Left)
			}
			if v.Right != nil {
				tmpQueue = append(tmpQueue, v.Right)
			}
		}
		queue = tmpQueue
		res++
	}

	return res
}

func main() {
	fmt.Println(maxDepth(nil))
}
