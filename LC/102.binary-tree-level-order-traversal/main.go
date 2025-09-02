package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			currentLevel[i] = node.Val
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel)
	}
	return result
}
