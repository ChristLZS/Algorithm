package main

import "fmt"

// 做到极致，不用库函数和额外空间
// 1. 先遍历一遍字符串，计算出空格的个数，然后扩容字符串
// 2. 从后往前遍历，遇到空格就替换成 %20，遇到非空格就从前拷贝到后，直到遍历完
func replaceSpace(s string) string {
	count := 0
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}

	res := make([]byte, 0, len(s)+count*2)
	res = append(res, s...)
	res = append(res, make([]byte, count*2)...)
	left := len(s) - 1
	for right := len(res) - 1; right >= 0; {
		if res[left] == ' ' {
			res[right] = '0'
			res[right-1] = '2'
			res[right-2] = '%'
			right -= 3
		} else {
			res[right] = res[left]
			right--
		}
		left--
	}

	return string(res)
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}
