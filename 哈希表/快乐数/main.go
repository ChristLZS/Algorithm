package main

func isHappy(n int) bool {
	m := make(map[int]struct{})
	n = getSum(n)
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		} else {
			m[n] = struct{}{}
		}
		n = getSum(n)
	}

	return true
}

func getSum(n int) int {
	sum := 0
	for t := n % 10; n != 0; t = n % 10 {
		sum += t * t
		n /= 10
	}
	return sum
}

func main() {
	result := isHappy(2)
	println(result)
}
