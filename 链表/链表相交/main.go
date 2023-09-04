package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路
// 长度对齐，让长链表移动至对齐位置，然后两个链表同时前进
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	lenA, lenB := 0, 0
	for node := headA; node != nil; node = node.Next {
		lenA++
	}
	for node := headB; node != nil; node = node.Next {
		lenB++
	}
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else if lenA < lenB {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
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
