package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 { // 递归终止条件：中序遍历为空
		return nil
	}

	// 寻找根节点
	val := preorder[0]

	// 寻找分割点
	var index int
	for i, v := range inorder {
		if v == val {
			index = i
			break
		}
	}

	return &TreeNode{
		Val:   val,
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
}

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
