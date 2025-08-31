package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		// 终止条件
		if node == nil {
			return
		}
		// 后续遍历
		traversal(node.Left)
		traversal(node.Right)
		// 记录结果
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}
