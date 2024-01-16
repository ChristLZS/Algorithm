package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := [][]int{}
	sort.Ints(nums)
	used := make([]int, len(nums)) //用于去重
	var dfs func([]int)
	dfs = func(tmpRes []int) {
		// 递归退出条件
		if len(tmpRes) == len(nums) {
			res = append(res, append([]int{}, tmpRes...))
			return
		}

		// 递归
		for i := 0; i < len(nums); i++ {
			// 去重：同一层递归中，如果前一个元素和当前元素相同，且前一个元素已经被使用过了，那么当前元素就不能使用
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
				continue
			}

			if used[i] == 0 { // == 0 表示没有被使用过
				used[i] = 1
				tmpRes = append(tmpRes, nums[i])
				dfs(tmpRes)
				tmpRes = tmpRes[:len(tmpRes)-1]
				used[i] = 0
			}
		}
	}

	dfs([]int{})

	return res
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
