package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	// 虚拟头节点
	dummyNode := &ListNode{Next: head}

	// fast 走 n 次
	index := 0
	slow := dummyNode
	fast := dummyNode
	for index < n {
		fast = fast.Next
		index++
		if fast == nil { // 输入N溢出，不处理
			return dummyNode.Next
		}
	}

	// fast先走一步，然后 slow 和 fast 同时走，直至 fast 到尾部，slow 做删除操作即可
	fast = fast.Next
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	// node := &ListNode{Val: 1}
	newNode := removeNthFromEnd(node, 1)
	for newNode != nil {
		fmt.Println(newNode.Val)
		newNode = newNode.Next
	}
}
