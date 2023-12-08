package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	curSum := 0
	res := [][]int{}
	tmpRes := []int{}
	var vist func(node *TreeNode)
	vist = func(node *TreeNode) {
		// cur
		if node != nil {
			curSum += node.Val // 计算当前和
			tmpRes = append(tmpRes, node.Val)
		}

		if node.Left == nil && node.Right == nil {
			if curSum == targetSum {
				r := make([]int, len(tmpRes))
				copy(r, tmpRes)
				res = append(res, r)
				return
			}
		}

		// left
		if node.Left != nil {
			vist(node.Left)
			curSum -= node.Left.Val // 回溯
			tmpRes = tmpRes[:len(tmpRes)-1]
		}

		// right
		if node.Right != nil {
			vist(node.Right)
			curSum -= node.Right.Val // 回溯
			tmpRes = tmpRes[:len(tmpRes)-1]
		}
	}

	vist(root)

	return res
}

func main() {
	fmt.Println(pathSum(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 5))
}
