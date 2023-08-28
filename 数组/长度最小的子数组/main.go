package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	left, right, sum, res := 0, 0, 0, len(nums)+1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < res {
				res = right - left + 1
			}
			sum -= nums[left]
			left++
		}

		right++
	}

	if res != len(nums)+1 {
		return res
	}

	return 0
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
