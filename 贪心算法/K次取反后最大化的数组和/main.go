package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	// 找到绝对值最小的数
	minAbs := nums[0]
	if minAbs < 0 {
		minAbs = -minAbs
	}

	minAbsIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if -nums[i] < minAbs {
				minAbs = -nums[i]
				minAbsIndex = i
			}
		} else {
			if nums[i] < minAbs {
				minAbs = nums[i]
				minAbsIndex = i
			}
		}
	}

	// 1. 优先反转负数
	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	// 2. 如果还有剩余的反转次数，则反转绝对值最小的数
	for k > 0 {
		nums[minAbsIndex] = -nums[minAbsIndex]
		k--
	}

	// 3. 计算和
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{-2, 5, 0, 2, -2}, 3))
}
