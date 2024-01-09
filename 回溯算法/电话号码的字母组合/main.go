package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	res := []string{}

	var data = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var dfs func(int, string, []byte)
	dfs = func(index int, digits string, curRes []byte) {
		// 回溯退出条件
		if index == len(digits) {
			res = append(res, string(curRes))
			return
		}

		// 递归
		num := digits[index]
		for i := 0; i < len(data[num]); i++ {
			curRes = append(curRes, data[num][i])
			dfs(index+1, digits, curRes)
			curRes = curRes[:len(curRes)-1]
		}

	}

	dfs(0, digits, []byte{})

	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
