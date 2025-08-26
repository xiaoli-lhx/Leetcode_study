package main

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
