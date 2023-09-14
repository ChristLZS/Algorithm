package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			num0 := stack[len(stack)-2]
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, num0+num1)
		case "-":
			num0 := stack[len(stack)-2]
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, num0-num1)
		case "*":
			num0 := stack[len(stack)-2]
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, num0*num1)
		case "/":
			num0 := stack[len(stack)-2]
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, num0/num1)
		default:
			data, _ := strconv.Atoi(tokens[i])
			stack = append(stack, data)
		}
	}

	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
