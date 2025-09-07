package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// 递归
func defs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return defs(left.Left, right.Right) && defs(left.Right, right.Left)
}
func isSymmetric(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}
