package main

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	n := len(haystack)
	m := len(needle)
	for i := 0; i <= n-m; i++ {
		j := 0
		for ; j < m; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == m {
			return i
		}
	}
	return -1
}

// TODO: 理解 KMP 算法

//func strStrKMP(haystack string, needle string) int {
//	n, m := len(haystack), len(needle)
//	//if m == 0 {
//	//	return 0
//	//}
//	if n < m {
//		return -1
//	}
//
//	// 1. 构建 next 数组
//	next := make([]int, m)
//	// j 是前缀的末尾，也是最长相等前后缀的长度
//	// i 是后缀的末尾
//	for i, j := 1, 0; i < m; i++ {
//		// 如果 j > 0 并且前后缀末尾字符不匹配，
//		// 则 j 回溯到上一个最长相等前后缀的位置
//		for j > 0 && needle[i] != needle[j] {
//			j = next[j-1]
//		}
//		// 如果前后缀末尾字符匹配，最长相等前后缀长度 +1
//		if needle[i] == needle[j] {
//			j++
//		}
//		next[i] = j
//	}
//
//	// 2. 使用 next 数组进行匹配
//	// j 是 needle 的指针，i 是 haystack 的指针
//	for i, j := 0, 0; i < n; i++ {
//		// 如果字符不匹配，并且 j > 0，
//		// j 就根据 next 数组回溯，i 保持不变
//		for j > 0 && haystack[i] != needle[j] {
//			j = next[j-1]
//		}
//		// 如果字符匹配，j 向后移动
//		if haystack[i] == needle[j] {
//			j++
//		}
//		// 如果 j 等于 needle 的长度，说明完全匹配
//		if j == m {
//			// 返回匹配的起始位置
//			return i - m + 1
//		}
//	}
//
//	return -1
//}
