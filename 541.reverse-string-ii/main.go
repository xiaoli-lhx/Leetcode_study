package main

// reverseStringHelper 是一个辅助函数，用于反转字节数组 s 的子切片。
// 从 left 索引到 right 索引的字符将被原地反转。
func reverseStringHelper(s []byte, left, right int) {
	// 使用双指针法原地反转
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// reverseStr 函数根据给定规则反转字符串 s。
//
// 规则：每 2k 个字符，反转前 k 个字符。
// 边缘情况:
//   - 剩余字符少于 k 个: 全部反转。
//   - 剩余字符少于 2k 但大于或等于 k 个: 只反转前 k 个。
//
// 参数:
//
//	s: 原始字符串
//	k: 整数，用于定义反转的长度
//
// 返回值:
//
//	反转后的新字符串。
func reverseStr(s string, k int) string {
	// 将字符串转换为字节数组，以便进行原地修改。
	// 这比创建新字符串更有效率。
	bytes := []byte(s)
	n := len(bytes)

	// 遍历字节数组，步长为 2k，处理每个子段。
	for i := 0; i < n; i += 2 * k {
		// 计算反转段的起始和结束索引。
		// 起始索引是当前的 i。
		left := i

		// 结束索引是 i + k - 1。
		// 但是我们需要考虑两种边缘情况：
		// 1. 如果剩余字符少于 k 个，right 应该是数组的最后一个索引 n - 1。
		// 2. 否则，right 应该是 i + k - 1。
		// 我们可以通过 min(i + k - 1, n - 1) 来统一处理。
		right := i + k - 1
		if right >= n {
			right = n - 1
		}

		// 调用辅助函数，反转指定范围内的字符。
		reverseStringHelper(bytes, left, right)
	}

	// 将修改后的字节数组转换回字符串并返回。
	return string(bytes)
}
