package main

import "fmt"

func reverseWords(s string) string {
	data := []byte(s)
	// 去除首尾空格
	for i := 0; i < len(data); i++ {
		if data[i] != ' ' {
			data = data[i:]
			break
		}
	}
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != ' ' {
			data = data[:i+1]
			break
		}
	}

	// 去除中间多余空格
	var count = 0
	var left, right = 0, 0
	for right < len(data) {
		if data[right] == ' ' {
			data[left] = ' '
			left++
			right++
			for data[right] == ' ' {
				right++
				count++
			}
		} else {
			data[left] = data[right]
			left++
			right++
		}
	}
	data = data[:len(data)-count]

	// 整体反转
	reverseString(data)

	// 单词反转
	left, right = 0, 0
	for right < len(data) {
		if data[right] == ' ' {
			reverseString(data[left:right])
			left = right + 1
		}
		right++
	}

	// 最后一个单词
	reverseString(data[left:])

	return string(data)
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
	fmt.Println(reverseWords("a good   example"))
}
