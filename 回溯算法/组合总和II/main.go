package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	res := [][]int{}
	tmp := []int{}
	used := make([]bool, len(candidates))
	sort.Ints(candidates)
	var dfs func(int, int, []int)
	dfs = func(index, sum int, tmp []int) {
		// 剪枝1: sum > target
		if sum > target {
			return
		}

		// 递归终止条件
		if sum == target {
			res = append(res, append([]int{}, tmp...))
			return
		}

		// 递归过程
		for i := index; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			} else {
				used[i] = true
			}

			tmp = append(tmp, candidates[i])
			dfs(i+1, sum+candidates[i], tmp)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}

	dfs(0, 0, tmp)

	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
