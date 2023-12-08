package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 用哈希表存储中序遍历的每个元素及其下标，方便查找根节点的下标
	m = make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}

	return build(inorder, postorder, len(postorder)-1, 0, len(inorder)-1)
}

// root: 后序遍历的根节点下标
// left: 中序遍历的左边界
// right: 中序遍历的右边界
// 返回值：构造出的二叉树根节点
func build(inorder, postorder []int, root, left, right int) *TreeNode {
	if left > right { // 递归终止条件：中序遍历为空
		return nil
	} else if left == right { // 递归终止条件：中序遍历只有一个元素
		return &TreeNode{Val: inorder[left]}
	}

	// 构造根节点
	node := &TreeNode{Val: postorder[root]}
	node.Left = build(inorder, postorder, root-(right-m[postorder[root]])-1, left, m[postorder[root]]-1)
	node.Right = build(inorder, postorder, root-1, m[postorder[root]]+1, right)
	return node
}

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
