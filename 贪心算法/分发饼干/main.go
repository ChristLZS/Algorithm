package main

import (
	"fmt"
	"sort"
)

func findContentChildren(child []int, food []int) int {
	// 1. 排序
	res := 0
	sort.Ints(child)
	sort.Ints(food)

	// 2. 遍历
	// 2.1 如果当前食物满足当前小孩，那么小孩和食物都往后移动一位
	// 2.2 如果当前食物不满足当前小孩，那么食物往后移动一位
	// 2.3 直到小孩或者食物遍历完
	for chIndex, foodIndex := 0, 0; chIndex < len(child) && foodIndex < len(food); {
		if child[chIndex] <= food[foodIndex] {
			res++
			chIndex++
			foodIndex++
		} else {
			foodIndex++
		}
	}

	return res
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}
