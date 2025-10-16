package main

//func productExceptSelf(nums []int) []int {
//	n := len(nums)
//	res := make([]int, n)
//	// 存放前缀积，由于第一个元素的前缀积为1，所以初始化为1
//	res[0] = 1
//	for i := 1; i < n; i++ {
//		// 前缀积分解为：res[i-1] * nums[i-1]
//		res[i] = res[i-1] * nums[i-1]
//	}
//	// 存放后缀积，由于最后一个元素的后缀积为1，所以初始化为1
//	right := 1
//	for i := n - 1; i >= 0; i-- {
//		// 结果分解为：res[i] * right
//		res[i] *= right
//		// 后缀积分解为：right * nums[i]
//		right *= nums[i]
//	}
//	return res
//}

// 更易读的写法
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	right := 1
	for i := len(nums) - 2; i >= 0; i-- {
		right *= nums[i+1]
		res[i] *= right
	}
	return res
}
