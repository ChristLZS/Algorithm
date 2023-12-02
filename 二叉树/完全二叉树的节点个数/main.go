package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if nil == root {
		return 0
	}

	res := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tmpQueue := []*TreeNode{}
		for _, node := range queue {
			res++
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

}
