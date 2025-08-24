package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	// 数组第一个元素肯定是唯一的
	slow := 1
	for fast := 1; fast < n; fast++ {
		// 如果当前元素和上一个元素不相同，则将当前元素复制到慢指针位置
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
