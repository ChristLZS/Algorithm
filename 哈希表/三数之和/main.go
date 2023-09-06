package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	// 去重
	m := make(map[string]bool)
	for _, v := range result {
		m[fmt.Sprintf("%d,%d,%d", v[0], v[1], v[2])] = true
	}
	result = [][]int{}
	for k := range m {
		var v [3]int
		fmt.Sscanf(k, "%d,%d,%d", &v[0], &v[1], &v[2])
		result = append(result, v[:])
	}

	return result
}

func main() {
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for _, v := range result {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
}
