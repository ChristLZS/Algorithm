package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	// 一个元素
	if len(nums) < 2 {
		return len(nums)
	}

	// 两个元素
	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return 1
		} else {
			return 2
		}
	}

	// 遍历数组
	var res = 1
	var preDiff = 0
	var curDiff = 0
	for i := 0; i < len(nums)-1; i++ {
		// 当前差值
		curDiff = nums[i+1] - nums[i]

		// 如果当前差值和上一个差值异号，那么就是一个峰值
		if preDiff >= 0 && curDiff < 0 || preDiff <= 0 && curDiff > 0 {
			res++
			preDiff = curDiff
		}
	}

	return res
}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
}
