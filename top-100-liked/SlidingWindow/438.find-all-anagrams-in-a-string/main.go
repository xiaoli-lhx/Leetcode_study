package main

func findAnagrams(s string, p string) []int {
	// 创建一个切片来存储结果
	var result []int
	n := len(s)
	m := len(p)

	// 边界情况：如果 s 比 p 短，不可能找到异位词
	if n < m {
		return result
	}

	// 用两个长度为 26 的数组作为频率计数器
	// pCount 用于存储 p 的字符频率
	// sCount 用于存储 s 中滑动窗口的字符频率
	var pCount [26]int
	var sCount [26]int

	// --- 1. 初始化阶段 ---
	// 统计 p 的频率，和 s 中第一个窗口的频率
	for i := 0; i < m; i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	// --- 2. 检查第一个窗口 ---
	// 如果第一个窗口就匹配，将起始索引 0 加入结果
	if pCount == sCount {
		result = append(result, 0)
	}

	// --- 3. 滑动窗口阶段 ---
	// 从索引 m 开始遍历 s，i 代表窗口的右边界
	for i := m; i < n; i++ {
		// 更新 sCount：
		// 新进入窗口的字符 s[i] 的计数加一
		sCount[s[i]-'a']++
		// 刚离开窗口的字符 s[i-m] 的计数减一
		sCount[s[i-m]-'a']--

		// --- 4. 检查滑动后的窗口 ---
		// 如果当前窗口的频率与 p 相同
		if pCount == sCount {
			// 计算并记录当前窗口的起始索引
			startIndex := i - m + 1
			result = append(result, startIndex)
		}
	}

	// 返回最终结果
	return result
}
