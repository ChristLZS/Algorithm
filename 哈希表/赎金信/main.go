package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		if m[v] > 0 {
			m[v]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	result := canConstruct("ab", "ba")
	println(result)
}
