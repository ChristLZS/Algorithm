package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := root.Val
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpQueue := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}

			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}

		if len(tmpQueue) > 0 {
			res = tmpQueue[0].Val
		}

		queue = tmpQueue
	}

	return res
}

func main() {
	fmt.Println(findBottomLeftValue(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
}
