package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}

	var dfs func(int, int, []int)
	dfs = func(index, curSum int, curRes []int) {
		// 回溯退出条件: 元素数量达标,且和为n
		if curSum == n && len(curRes) == k {
			res = append(res, append([]int{}, curRes...))
			return
		}

		// 递归
		for i := index; i <= 9; i++ {
			// 剪枝1：和超过了
			if curSum+i > n {
				return
			}

			// 剪枝2：剩余的元素个数不够了
			if len(curRes)+(9-i+1) < k {
				return
			}

			curSum += i
			curRes = append(curRes, i)
			dfs(i+1, curSum, curRes)
			curRes = curRes[:len(curRes)-1]
			curSum -= i
		}
	}
	dfs(1, 0, []int{})

	return res
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}
