package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	ans := [][]int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			level = append(level, node.Val)
			queue = queue[1:]
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		ans = append(ans, level)
	}
	return ans
}
