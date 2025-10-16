package main

// 前缀和 + 哈希表
// 关键点：
// 1. 前缀和数组
// 2. 哈希表存储前缀和出现的次数
// 目标求解： sum(i,j) = k
// sum(i, j) = prefix_sum[j] - prefix_sum[i-1]
// 则有： prefix_sum[i-1] = prefix_sum[j] - k

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range nums {
		sum += num
		if _, ok := m[sum-k]; ok {
			count++
		}
		m[sum]++
	}
	return count
}
