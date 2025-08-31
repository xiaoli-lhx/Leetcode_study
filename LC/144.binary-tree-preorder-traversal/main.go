package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	func preorderTraversal(root *TreeNode) []int {
//		var res []int
//		var traversal func(node *TreeNode)
//		traversal = func(node *TreeNode) {
//			// 终止条件
//			if node == nil {
//				return
//			}
//			// 先序遍历
//			res = append(res, node.Val)
//			// 左子树
//			traversal(node.Left)
//			// 右子树
//			traversal(node.Right)
//		}
//		traversal(root)
//		return res
//	}
//
// 迭代法
func preorderTraversal(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return res
}
