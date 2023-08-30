package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return nil
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	newNode := reverseList(node)
	for newNode != nil {
		println(newNode.Val)
		newNode = newNode.Next
	}
}
