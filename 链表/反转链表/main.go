package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 递归
func reverseList2(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return reverseList2(cur, next)
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	// newNode := reverseList(node)
	newNode := reverseList2(nil, node)
	for newNode != nil {
		fmt.Println(newNode.Val)
		newNode = newNode.Next
	}
}
