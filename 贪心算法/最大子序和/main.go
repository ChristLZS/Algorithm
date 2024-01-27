package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 贪心算法
	// 从第一个元素开始，如果当前元素之前的和小于0，则丢弃之前的元素
	// 如果当前元素之前的和大于0，则加上当前元素
	// 每次加上当前元素后，都要判断当前和是否大于最大和
	// 如果大于最大和，则更新最大和
	// 最后返回最大和
	var maxSum, tmpSum int
	maxSum = nums[0]
	tmpSum = nums[0]
	for i := 1; i < len(nums); i++ {
		if tmpSum < 0 {
			tmpSum = 0
		}

		if tmpSum += nums[i]; tmpSum > maxSum {
			maxSum = tmpSum
		}

	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
