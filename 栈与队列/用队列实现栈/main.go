package main

import "fmt"

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{stackIn: []int{}, stackOut: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stackOut) == 0 { // stackOut为空，将stackIn中的元素全部转移到stackOut中
		for len(this.stackIn) > 0 {
			this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
		}
	}

	if len(this.stackOut) == 0 { // stackOut为空，返回-1
		return -1
	}

	// stackOut不为空，返回stackOut的最后一个元素
	res := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]

	return res
}

func (this *MyQueue) Peek() int {
	if len(this.stackOut) == 0 { // stackOut为空，将stackIn中的元素全部转移到stackOut中
		for len(this.stackIn) > 0 {
			this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
		}
	}

	if len(this.stackOut) == 0 { // stackOut为空，返回-1
		return -1
	}

	// stackOut不为空，返回stackOut的最后一个元素
	return this.stackOut[len(this.stackOut)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
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
