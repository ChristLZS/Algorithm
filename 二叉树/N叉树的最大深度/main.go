package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	queue := []*Node{root}
	for len(queue) > 0 {
		maxDepth++
		tmpQueue := []*Node{}
		for _, node := range queue {
			if node != nil {
				tmpQueue = append(tmpQueue, node.Children...)
			}
		}

		queue = tmpQueue
	}

	return maxDepth
}

func main() {
	// [1,null,3,2,4,null,5,6]
	fmt.Println(maxDepth(&Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 5,
					},
					&Node{
						Val: 6,
					},
				},
			},
			&Node{
				Val: 2,
			},
			&Node{
				Val: 4,
			},
		},
	}))
}
