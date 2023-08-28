package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	dummyHead *ListNode
	size      int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	return 0
}

func (this *MyLinkedList) AddAtHead(val int) {

}

func (this *MyLinkedList) AddAtTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

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
