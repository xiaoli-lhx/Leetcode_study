package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := m[target-nums[i]]; ok {
			return []int{j, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
