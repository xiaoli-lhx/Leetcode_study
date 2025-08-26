package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// lastSeen 哈希表：key 是字符，value 是该字符最近一次出现的索引
	lastSeen := make(map[rune]int)
	maxLength := 0
	// left 是滑动窗口的左边界
	left := 0

	// right 是滑动窗口的右边界
	for right, char := range s {
		// 检查当前字符 char 是否在哈希表中，并且其上次出现的位置是否在当前窗口内
		if lastIndex, found := lastSeen[char]; found && lastIndex >= left {
			// 如果是，说明遇到了重复字符。
			// 我们需要收缩窗口，将左边界移动到重复字符上次出现位置的下一个位置。
			left = lastIndex + 1
		}

		// 计算当前窗口的长度
		currentLength := right - left + 1
		// 更新最大长度
		if currentLength > maxLength {
			maxLength = currentLength
		}

		// 无论如何，都要更新当前字符的最新位置
		lastSeen[char] = right
	}

	return maxLength
}

func main() {
	// 测试用例
	str1 := "abcabcbb"
	fmt.Printf("'%s' 的最长无重复子串长度是: %d\n", str1, lengthOfLongestSubstring(str1)) // 期望输出: 3 ("abc")

	str2 := "bbbbb"
	fmt.Printf("'%s' 的最长无重复子串长度是: %d\n", str2, lengthOfLongestSubstring(str2)) // 期望输出: 1 ("b")

	str3 := "pwwkew"
	fmt.Printf("'%s' 的最长无重复子串长度是: %d\n", str3, lengthOfLongestSubstring(str3)) // 期望输出: 3 ("wke")

	str4 := "arabcacfr"
	fmt.Printf("'%s' 的最长无重复子串长度是: %d\n", str4, lengthOfLongestSubstring(str4)) // 期望输出: 4 ("rabc" and "acfr")
}
