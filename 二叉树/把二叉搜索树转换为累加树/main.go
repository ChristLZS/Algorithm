package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 倒序中序遍历，使用一个变量记录前一个节点，然后当前节点加上前一个节点的值
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var prev *TreeNode
	var visit func(*TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}

		visit(node.Right)

		if prev != nil {
			node.Val += prev.Val
		}
		prev = node

		visit(node.Left)
	}

	visit(root)

	return root
}

func main() {
	// [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
	fmt.Println(convertBST(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  7,
				Left: nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}))
}
