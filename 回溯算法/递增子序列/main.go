package main

import (
	"fmt"
)

func findSubsequences(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var res [][]int
	var dfs func(int, []int)
	dfs = func(index int, tmp []int) {
		used := make(map[int]bool)

		// 递归
		for i := index; i < len(nums); i++ {
			// 剪枝1：去重
			if used[nums[i]] {
				continue
			} else {
				used[nums[i]] = true
			}

			// 剪枝2：保证递增
			if len(tmp) > 0 && nums[i] < tmp[len(tmp)-1] {
				continue
			}

			tmp = append(tmp, nums[i])
			if len(tmp) > 1 {
				res = append(res, append([]int{}, tmp...))
			}
			dfs(i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0, []int{})
	return res
}

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}
