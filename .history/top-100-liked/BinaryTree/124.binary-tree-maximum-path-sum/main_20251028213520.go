package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		// 递归终止条件
		if node == nil {
			return 0
		}
		leftGain := max(0, maxGain(node.Left))
		rightGain := max(0, maxGain(node.Right))
		pathAtThisNode := node.Val + leftGain + rightGain
		maxSum = max(maxSum, pathAtThisNode)
		return node.Val + max(rightGain, leftGain)
	}
	maxGain(root)
	return maxSum
}
