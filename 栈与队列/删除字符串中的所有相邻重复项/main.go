package main

import "fmt"

func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] { // 栈顶元素与当前元素相同
			stack = stack[:len(stack)-1] // 出栈
		} else {
			stack = append(stack, s[i]) // 入栈
		}
	}
	return string(stack)
}

func main() {
	s := "abbaca"
	fmt.Println(removeDuplicates(s))
}
