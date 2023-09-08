package main

import "fmt"

func reverseStr(s string, k int) string {
	res := []byte(s)
	for i := 0; i < len(res); i += 2 * k {
		if i+k <= len(res) {
			reverseString(res[i : i+k])
		} else {
			reverseString(res[i:])
		}
	}
	return string(res)
}

func reverseString(s []byte) {
	var left, right = 0, len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}
