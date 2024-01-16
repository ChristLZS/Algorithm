package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := [][]int{}
	used := make(map[int]bool)
	var dfs func([]int)
	dfs = func(tmpRes []int) {
		// 递归退出条件
		if len(tmpRes) == len(nums) {
			res = append(res, append([]int{}, tmpRes...))
			return
		}

		// 递归
		for i := 0; i < len(nums); i++ { // 全排列，不需要剪枝
			if !used[i] {
				used[i] = true
				tmpRes = append(tmpRes, nums[i])
				dfs(tmpRes)
				tmpRes = tmpRes[:len(tmpRes)-1]
				used[i] = false
			}
		}
	}

	dfs([]int{})

	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
