package main

import "fmt"

// 动态规划解法
//func trap(height []int) int {
//	n := len(height)
//	leftMax := make([]int, n)
//	rightMax := make([]int, n)
//	leftMax[0] = height[0]
//	for i := 1; i < n; i++ {
//		leftMax[i] = max(leftMax[i-1], height[i])
//		// fmt.Println("leftMax[i]", leftMax[i])
//	}
//	rightMax[n-1] = height[n-1]
//	for i := n - 2; i >= 0; i-- {
//		rightMax[i] = max(rightMax[i+1], height[i])
//	}
//	var water int
//	for i := 1; i < n-1; i++ {
//		water += min(leftMax[i], rightMax[i]) - height[i]
//	}
//	return water
//}

// 双指针解法
func trap(height []int) int {
	n := len(height)
	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	totalWater := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
			right--
		}
	}
	return totalWater
}

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}
