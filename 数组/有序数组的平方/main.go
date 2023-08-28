package main

import "fmt"

func sortedSquares(nums []int) []int {
	index, left, right := len(nums)-1, 0, len(nums)-1
	res := make([]int, len(nums))
	for left <= right {
		l := nums[left] * nums[left]
		r := nums[right] * nums[right]
		if l <= r {
			res[index] = r
			right--
		} else {
			res[index] = l
			left++
		}
		index--
	}

	return res
}

func main() {
	fmt.Printf("%+#v", sortedSquares([]int{-4, -1, 0, 3, 10}))
}
