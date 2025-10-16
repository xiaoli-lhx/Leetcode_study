package main

// 使用一个新的数组
//func rotate(nums []int, k int) {
//	k = k % len(nums)
//	newNums := make([]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		newIndex := (i + k) % len(nums)
//		newNums[newIndex] = nums[i]
//	}
//	copy(nums, newNums)
//}

// 三次翻转
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, k int, i int) {
	for k < i {
		nums[k], nums[i] = nums[i], nums[k]
		k++
		i--
	}
}
