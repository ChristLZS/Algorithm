package main

func threeSum(nums []int) [][]int {
	return nil
}

func main() {
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for _, v := range result {
		for _, v2 := range v {
			println(v2)
		}
	}
}
