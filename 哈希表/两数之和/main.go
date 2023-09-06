package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		if j, ok := m[target-v]; ok && i != j {
			return []int{i, j}
		}
	}

	return nil
}

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	for _, v := range result {
		println(v)
	}
}
