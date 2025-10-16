package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		count++
	}
	return count
}

// 递归 DFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}
