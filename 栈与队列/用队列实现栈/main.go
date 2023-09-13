package main

import "fmt"

type MyStack struct {
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {

}

func (this *MyStack) Pop() int {
	return 0
}

func (this *MyStack) Top() int {
	return 0
}

func (this *MyStack) Empty() bool {
	return false
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
