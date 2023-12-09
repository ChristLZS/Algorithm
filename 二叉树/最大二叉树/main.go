package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 { // 递归终止条件
		return nil
	}

	max := 0
	index := -1
	for i, v := range nums {
		if v > nums[max] {
			max = i
			index = i
		}
	}

	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}
