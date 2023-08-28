package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	fakeHead := &ListNode{}
	fakeHead.Next = head
	cur := fakeHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return fakeHead.Next
}

func main() {
	// 1 2 2 1
	// 1
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	res := removeElements(head, 2)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}
