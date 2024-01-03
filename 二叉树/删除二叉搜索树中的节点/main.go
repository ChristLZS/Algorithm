package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	var parent = root
	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil { // case 0: 没找到
			return
		}

		if key < node.Val { // 左子树
			parent = node
			visit(node.Left)
			return
		} else if key > node.Val { // 右子树
			parent = node
			visit(node.Right)
			return
		} else { // 找到了
			// case 1: 叶子节点
			if node.Left == nil && node.Right == nil {
				if parent.Left == node {
					parent.Left = nil
				} else if parent.Right == node {
					parent.Right = nil
				} else { // 要删除的是根节点
					root = nil
				}
				return
			}

			// case 2: 只有左子树
			if node.Left != nil && node.Right == nil {
				if parent.Left == node {
					parent.Left = node.Left
				} else if parent.Right == node {
					parent.Right = node.Left
				} else {
					root = node.Left // 要删除的是根节点
				}
				return
			}

			// case 3: 只有右子树
			if node.Left == nil && node.Right != nil {
				if parent.Left == node {
					parent.Left = node.Right
				} else if parent.Right == node {
					parent.Right = node.Right
				} else {
					root = node.Right // 要删除的是根节点
				}
				return
			}

			// case 4: 左右子树都有
			if node.Left != nil && node.Right != nil {
				if parent.Right == node {
					// 找到右子树最小的节点
					var min = node.Right
					for min.Left != nil {
						min = min.Left
					}

					// 最小节点的左子树指向当前节点的左子树
					min.Left = node.Left

					// 删除该节点
					parent.Right = node.Right
				} else if parent.Left == node {
					// 找到左子树最大的节点
					var max = node.Left
					for max.Right != nil {
						max = max.Right
					}

					// 最大节点的右子树指向当前节点的右子树
					max.Right = node.Right

					// 删除该节点
					parent.Left = node.Left
				} else { // 要删除的是根节点
					// 找到左子树最大的节点
					var max = node.Left
					for max.Right != nil {
						max = max.Right
					}

					// 最大节点的右子树指向当前节点的右子树
					max.Right = node.Right

					// 删除该节点
					root = node.Left
				}
			}
		}
	}

	visit(root)
	return root
}

func main() {
	// [5,3,6,2,4,null,7]
	// 3
	fmt.Println(deleteNode(&TreeNode{Val: 0}, 0))
}
