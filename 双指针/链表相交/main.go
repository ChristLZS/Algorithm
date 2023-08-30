package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return nil
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	newNode := getIntersectionNode(node, node)
	for newNode != nil {
		println(newNode.Val)
		newNode = newNode.Next
	}
}
