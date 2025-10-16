package main

// []int{0, 1, 0, 3, 12} -> []int{1, 3, 12, 0, 0}
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
