package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := [][]int{}
	res = append(res, []int{})
	sort.Ints(nums)
	var dfs func(int, []int)
	dfs = func(index int, tmp []int) {
		// 递归
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}

			tmp = append(tmp, nums[i])
			res = append(res, append([]int{}, tmp...))
			dfs(i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(0, []int{})
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
