package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	return nil
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	newNode := detectCycle(node)
	for newNode != nil {
		println(newNode.Val)
		newNode = newNode.Next
	}
}
