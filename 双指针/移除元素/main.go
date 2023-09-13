package main

func removeElement(nums []int, val int) int {
	var slow, fast = 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	result := removeElement([]int{3, 2, 2, 3}, 3)
	println(result)
}
