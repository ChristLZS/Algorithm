package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		tmpQueue := make([]*Node, 0)
		tmpRes := make([]int, 0)
		for _, v := range queue {
			tmpRes = append(tmpRes, v.Val)
			tmpQueue = append(tmpQueue, v.Children...)
		}
		queue = tmpQueue
		res = append(res, tmpRes)
	}

	return res
}

func main() {
	fmt.Println(levelOrder(nil))
}
