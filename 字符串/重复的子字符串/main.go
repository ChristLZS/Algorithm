package main

import "fmt"

// 思路：KMP算法
// 1.求next数组
// 2.判断next[n-1]+1是否为最长相同前后缀的长度，如果是，判断n%(n-(next[n-1]+1))是否为0，如果是，返回true
// 3.否则返回false
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	next := make([]int, n)
	next[0] = -1
	j := -1
	for i := 1; i < n; i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	// next[n-1]+1 最长相同前后缀的长度
	if next[n-1] != -1 && n%(n-(next[n-1]+1)) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
}
