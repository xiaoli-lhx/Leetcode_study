package main

func canConstruct(ransomNote string, magazine string) bool {
	// 如果ransomNote的长度大于magazine的长度，则肯定不可能组成
	if len(ransomNote) > len(magazine) {
		return false
	}
	charCount := make(map[rune]int)
	for _, c := range magazine {
		charCount[c]++
	}
	for _, c := range ransomNote {
		if charCount[c] > 0 {
			charCount[c]--
		} else {
			return false
		}
	}
	return true
}
