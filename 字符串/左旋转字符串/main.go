package main

import "fmt"

// 思路1：用一个额外的byte数组，每次将前1个字符放到数组后面，再将后面的字符放到数组前面
// 思路2：翻转前n个字符，翻转后面的字符，最后翻转整个字符串
func reverseLeftWords(s string, n int) string {
	data := []byte(s)
	for i := 0; i < n; i++ {
		s := data[0]
		for j := 0; j < len(data)-1; j++ {
			data[j] = data[j+1]
		}
		data[len(data)-1] = s
	}
	return string(data)
}

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}
