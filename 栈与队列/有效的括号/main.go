package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' { // 左括号入栈
			stack = append(stack, s[i])
		} else if len(stack) == 0 { // 右括号，但栈为空
			return false
		} else if s[i] == ')' && stack[len(stack)-1] != '(' { // 右括号，但栈顶不是对应的左括号
			return false
		} else if s[i] == ']' && stack[len(stack)-1] != '[' { // 右括号，但栈顶不是对应的左括号
			return false
		} else if s[i] == '}' && stack[len(stack)-1] != '{' { // 右括号，但栈顶不是对应的左括号
			return false
		} else { // 右括号，且栈顶是对应的左括号
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 { // 栈不为空
		return false
	}

	return true
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
	s = "([)]"
	fmt.Println(isValid(s))
}
