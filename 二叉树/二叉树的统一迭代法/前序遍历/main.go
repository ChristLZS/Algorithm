package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 统一迭代法
// 前序遍历：根左右
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	stack = append(stack, cur)
	for len(stack) > 0 {
		// 处理一个节点
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil { // 添加节点
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			stack = append(stack, cur)
			stack = append(stack, nil) // 标记节点
		} else { // 处理节点
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
		}
	}

	return res
}

func main() {
	preorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}})
}
