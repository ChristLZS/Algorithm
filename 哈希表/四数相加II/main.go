package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	result := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			result += m[-v3-v4]
		}
	}
	return result
}

func main() {
	result := fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
	println(result)
}
