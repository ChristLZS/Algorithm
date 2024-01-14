package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	res := [][]int{}
	res = append(res, []int{})
	var dfs func(int, []int)
	dfs = func(index int, tmp []int) {
		// 递归
		for i := index; i < len(nums); i++ {
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
	fmt.Println(subsets([]int{1, 2, 3}))
}
