package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return nil
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	newNode := removeNthFromEnd(node, 2)
	for newNode != nil {
		println(newNode.Val)
		newNode = newNode.Next
	}
}
