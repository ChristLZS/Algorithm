package main

func fourSum(nums []int, target int) [][]int {
	return nil
}

func main() {
	result := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	for _, v := range result {
		for _, v2 := range v {
			println(v2)
		}
	}
}
