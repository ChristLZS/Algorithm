package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	return nil
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	newNode := swapPairs(node)
	for newNode != nil {
		fmt.Println(newNode.Val)
		newNode = newNode.Next
	}
}
