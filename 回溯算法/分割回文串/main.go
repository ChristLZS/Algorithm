package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	return s[0] == s[len(s)-1] && isPalindrome(s[1:len(s)-1])
}

func partition(s string) [][]string {
	res := [][]string{}
	tmpRes := []string{}
	var dfs func(string, int, []string)
	dfs = func(s string, index int, tmpRes []string) {
		// 递归终止条件
		if index == len(s) {
			res = append(res, append([]string{}, tmpRes...))
			return
		}

		// 递归过程
		for i := index; i < len(s); i++ {
			if isPalindrome(s[index : i+1]) {
				tmpRes = append(tmpRes, s[index:i+1])
				dfs(s, i+1, tmpRes)
				tmpRes = tmpRes[:len(tmpRes)-1]
			}
		}
	}

	dfs(s, 0, tmpRes)

	return res
}

func main() {
	fmt.Println(partition("aab"))
}
