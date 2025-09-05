package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		count++
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return count
			}
		}

	}
	return count
}
