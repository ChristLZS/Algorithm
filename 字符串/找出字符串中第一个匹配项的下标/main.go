package main

import "fmt"

// getNext 构造前缀表next
// params:
//
//	next 前缀表数组
//	s 模式串
func getNext(s string) []int {
	j := -1 // j表示 最长相等前后缀长度
	next := make([]int, len(s))
	next[0] = -1

	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j] // 回退前一位
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
	return next
}

// strStr KMP算法
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := getNext(needle)
	j := -1 // 模式串的起始位置 next为-1 因此也为-1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // 寻找下一个匹配点
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 { // j指向了模式串的末尾
			return i - len(needle) + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
