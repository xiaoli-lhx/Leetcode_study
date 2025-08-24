package main

//	func isAnagram(s string, t string) bool {
//		// 如果两个字符串长度不同，则肯定不是字母异位词
//		if len(s) != len(t) {
//			return false
//		}
//		count := [26]int{} // 定义一个计数数组
//		for i := 0; i < len(s); i++ {
//			count[s[i]-'a']++
//			count[t[i]-'a']--
//		}
//		for i := 0; i < 26; i++ {
//			if count[i] != 0 {
//				return false
//			}
//		}
//		return true
//	}
func isAnagram(s string, t string) bool {
	rs := []rune(s)
	rt := []rune(t)
	if len(rs) != len(rt) {
		return false
	}
	count := make(map[rune]int)
	for _, v := range rs {
		count[v]++
	}
	for _, v := range rt {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}
	return true
}
