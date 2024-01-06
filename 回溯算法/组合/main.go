package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}

	var dfs func(int, []int)
	dfs = func(index int, curRes []int) {
		// 回溯退出条件
		if len(curRes) == k {
			res = append(res, append([]int{}, curRes...))
			return
		}

		// 递归
		for i := index; i <= n; i++ {
			curRes = append(curRes, i)
			dfs(i+1, curRes)
			curRes = curRes[:len(curRes)-1]
		}
	}

	dfs(1, []int{})

	return res
}

func main() {
	// 4 2
	fmt.Println(combine(4, 2))
}
