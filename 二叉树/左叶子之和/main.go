package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func sumOfLeftLeaves(root *TreeNode) int {
	if nil == root {
		return 0
	}

	res := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpQueue := []*TreeNode{}
		for _, node := range queue {
			if left := node.Left; left != nil {
				// 判断是否为叶子节点
				if left.Left == nil && left.Right == nil {
					res += left.Val
				}

				tmpQueue = append(tmpQueue, left)
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
	fmt.Println(sumOfLeftLeaves(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
}
