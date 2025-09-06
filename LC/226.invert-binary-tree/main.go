package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序递归
//func invertTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	root.Left, root.Right = root.Right, root.Left
//	invertTree(root.Left)
//	invertTree(root.Right)
//	return root
//}

// 迭代
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}
