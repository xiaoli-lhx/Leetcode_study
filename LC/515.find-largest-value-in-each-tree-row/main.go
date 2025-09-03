package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		maxVal := queue[0].Val
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Val > maxVal {
				maxVal = node.Val
			}
		}
		ans = append(ans, maxVal)
	}
	return ans
}
