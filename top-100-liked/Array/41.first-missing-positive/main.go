package main

// 使用了哈希表
//func firstMissingPositive(nums []int) int {
//	hashMap := make(map[int]bool)
//	n := len(nums)
//	for i := 0; i < n; i++ {
//		hashMap[nums[i]] = true
//	}
//	// 因为数组长度为n，所以所需的最小的正整数，不可能大于n+1
//	for i := 1; i <= n+1; i++ {
//		if _, ok := hashMap[i]; !ok {
//			return i
//		}
//	}
//	return 1
//}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
