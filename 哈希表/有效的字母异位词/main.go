package main

// 思路
// 1. 将 s 中的字符放入 map 中，key 为字符，value 为出现的次数
// 2. 遍历 t 中的字符，如果字符在 map 中，则 value 减 1, 否则返回 false
// 3. 遍历 map，如果存在 value 不为 0，则返回 false, 否则返回 true
func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	for _, v := range s {
		sMap[v] = sMap[v] + 1
	}
	for _, v := range t {
		if _, ok := sMap[v]; ok {
			sMap[v] = sMap[v] - 1
		} else {
			return false
		}
	}
	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	println(isAnagram("anagram", "nagaram")) // true
	println(isAnagram("rat", "car"))         // false
}
