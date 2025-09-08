package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 1. 将所有数记录在哈希表中，自动处理掉重复项
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxCount := 0

	// 2. 直接遍历哈希表中的唯一数字
	for num := range numSet {
		// 3. 判断 num-1 是否在哈希表中
		// 如果在，就跳过，因为它不是起点
		if numSet[num-1] {
			continue
		}

		// 4. 如果 num-1 不在，说明 num 是一个序列的起点
		// 开始向后查找
		currentNum := num
		currentCount := 1

		for numSet[currentNum+1] {
			currentNum++
			currentCount++
		}

		// 5. 更新最大长度
		if currentCount > maxCount {
			maxCount = currentCount
		}
	}

	return maxCount
}
