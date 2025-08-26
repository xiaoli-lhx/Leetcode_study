package main

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	n := make(map[byte]int)
	m := make(map[byte]int)
	start := 0
	for i := 0; i < len(s); i++ {
		n[t[i]]++
	}
	match := 0
	minLen := len(s) + 1
	left, right := 0, 0
	for right < len(s) {
		c1 := s[right]
		right++
		if _, ok := n[c1]; ok {
			m[c1]++
			if m[c1] == n[c1] {
				match++
			}
		}
		for match == len(n) {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			c2 := s[left]
			left++
			if _, ok := n[c2]; ok {
				if m[c2] == n[c2] {
					match--
				}
				m[c2]--
			}
		}
	}
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
