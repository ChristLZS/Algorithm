package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			lefr := j + 1
			right := len(nums) - 1
			for lefr < right {
				sum := nums[i] + nums[j] + nums[lefr] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[lefr], nums[right]})
					right--
					lefr++
				} else if sum > target {
					right--
				} else {
					lefr++
				}
			}
		}
	}

	// 去重
	var m = make(map[string]bool)
	for _, v := range result {
		m[fmt.Sprintf("%d,%d,%d,%d", v[0], v[1], v[2], v[3])] = true
	}
	result = [][]int{}
	for k := range m {
		var v [4]int
		fmt.Sscanf(k, "%d,%d,%d,%d", &v[0], &v[1], &v[2], &v[3])
		result = append(result, v[:])
	}

	return result
}

func main() {
	result := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	for _, v := range result {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
}
