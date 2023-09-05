package main

func intersection(nums1 []int, nums2 []int) []int {
	res := make(map[int]struct{}, 0)
	tmp := make(map[int]struct{})
	for _, v := range nums1 {
		tmp[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := tmp[v]; ok {
			res[v] = struct{}{}
		}
	}

	result := make([]int, 0, len(res))
	for k := range res {
		result = append(result, k)
	}

	return result
}

func main() {
	result := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	for _, v := range result {
		println(v)
	}
}
