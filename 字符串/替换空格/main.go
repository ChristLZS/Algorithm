package main

import "fmt"

// 做到极致，不用库函数和额外空间
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
