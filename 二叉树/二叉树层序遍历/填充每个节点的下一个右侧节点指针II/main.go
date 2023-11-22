package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		tmpQueue := make([]*Node, 0)
		for index, v := range queue {
			if index+1 < len(queue) {
				v.Next = queue[index+1]
			} else {
				v.Next = nil
			}

			if v.Left != nil {
				tmpQueue = append(tmpQueue, v.Left)
			}
			if v.Right != nil {
				tmpQueue = append(tmpQueue, v.Right)
			}
		}
		queue = tmpQueue
	}

	return root
}

func main() {
	// [1,2,3,4,5,6,7]
	fmt.Println(connect(&Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}))
}
