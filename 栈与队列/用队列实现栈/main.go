package main

import "fmt"

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{data: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	if len(this.data) == 0 {
		return -1
	} else {
		data := this.data[0]
		this.data = this.data[1:]
		return data
	}
}

func (this *MyQueue) Peek() int {
	if len(this.data) == 0 {
		return -1
	}

	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	t := queue.Peek() // 返回 1
	fmt.Println(t)
	t = queue.Pop() // 返回 1
	fmt.Println(t)
	queue.Empty() // 返回 false
}
