package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func inorderTraversal(root *TreeNode) []int {
//	var res []int
//	var traversal func(node *TreeNode)
//	traversal = func(node *TreeNode) {
//		// 终止条件
//		if node == nil {
//			return
//		}
//		// 中序遍历
//		traversal(node.Left)
//		res = append(res, node.Val)
//		traversal(node.Right)
//	}
//	traversal(root)
//	return res
//}

// 迭代法
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
	}
	return res
}
