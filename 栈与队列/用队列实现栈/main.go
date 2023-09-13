package main

import (
	"fmt"
)

// 单队列实现栈
type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{queue: []int{}}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	if len(this.queue) == 0 {
		return -1
	}

	length := len(this.queue)
	for i := 0; i < length-1; i++ { // 将队列前面的元素依次出队再入队，直到队列中只剩下一个元素,这个元素就是栈顶元素
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
	}
	data := this.queue[0]
	this.queue = this.queue[1:]
	return data
}

func (this *MyStack) Top() int {
	if len(this.queue) == 0 {
		return -1
	}

	length := len(this.queue)
	for i := 0; i < length-1; i++ { // 将队列前面的元素依次出队再入队，直到队列中只剩下一个元素,这个元素就是栈顶元素
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
	}
	data := this.queue[0]
	this.queue = append(this.queue, this.queue[0]) // 将栈顶元素再次入队

	return data
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	t := stack.Top() // 返回 2
	fmt.Println(t)
	t = stack.Pop() // 返回 2
	fmt.Println(t)
	stack.Empty() // 返回 False
}
