package main

func maxSlidingWindow(nums []int, k int) []int {
	// 1. 准备工作
	// 用一个切片来做双端队列，存放索引
	deque := []int{}
	result := []int{}
	// 2. 遍历数组
	for i := 0; i < len(nums); i++ {
		// -清理队尾
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		// -加入队尾
		deque = append(deque, i)
		// -清理队头
		if deque[0] <= i-k {
			deque = deque[1:]
		}
		// -加入结果
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
