package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, sum, minlen := 0, 0, n+1
	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum >= target {
			count := right - left + 1
			if count < minlen {
				minlen = count
			}
			sum -= nums[left]
			left++
		}

	}
	if minlen == n+1 {
		return 0
	}
	return minlen
}
