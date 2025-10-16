package main

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
