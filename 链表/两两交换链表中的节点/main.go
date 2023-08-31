package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// before step 1
// /// pre      cur   next
// dummyNode --> 1 --> 2 --> 3 --> 4

// before step 2
// pre --> next --> cur --> 3 --> 4
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var dummyNode = &ListNode{Next: head}

	pre := dummyNode
	cur := dummyNode.Next
	next := cur.Next
	for cur != nil && next != nil {
		// step 1
		pre.Next = next
		tmp := next.Next
		next.Next = cur
		cur.Next = tmp

		// step 2
		pre = cur
		cur = pre.Next
		if cur != nil {
			next = cur.Next
		}
	}

	return dummyNode.Next
}

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	newNode := swapPairs(node)
	for newNode != nil {
		fmt.Println(newNode.Val)
		newNode = newNode.Next
	}
}
