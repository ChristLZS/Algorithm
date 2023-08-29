package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	dummyHead *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &Node{},
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	cur := l.dummyHead
	for i := 0; i <= index; i++ {
		if cur != nil {
			cur = cur.Next
		} else {
			return -1
		}
	}
	if cur != nil {
		return cur.Val
	} else {
		return -1
	}
}

func (l *MyLinkedList) AddAtHead(val int) {
	cur := l.dummyHead
	node := &Node{Val: val, Next: cur.Next}
	cur.Next = node
}

func (l *MyLinkedList) AddAtTail(val int) {
	cur := l.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &Node{Val: val}
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	cur := l.dummyHead
	for i := 0; i < index; i++ {
		if cur.Next != nil {
			cur = cur.Next
		} else {
			return
		}
	}
	node := &Node{Val: val, Next: cur.Next}
	cur.Next = node
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	cur := l.dummyHead
	for i := 0; i < index; i++ {
		if cur.Next != nil {
			cur = cur.Next
		}
	}

	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
}

func main() {
	// Your MyLinkedList object will be instantiated and called as such:
	obj := Constructor()
	param_1 := obj.Get(0)
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	obj.DeleteAtIndex(1)
	fmt.Println(param_1)
}
