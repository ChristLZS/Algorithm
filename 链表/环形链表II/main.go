package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路
// 1.快慢指针，判断是否为环型
// 2.如果是环型，快指针回到头节点，然后快慢指针同时前进，相遇的节点即为环的入口
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}

			return head
		}
	}
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
