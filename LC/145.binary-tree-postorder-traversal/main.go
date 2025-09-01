package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func postorderTraversal(root *TreeNode) []int {
//	var res []int
//	var traversal func(node *TreeNode)
//	traversal = func(node *TreeNode) {
//		// 终止条件
//		if node == nil {
//			return
//		}
//		// 后续遍历
//		traversal(node.Left)
//		traversal(node.Right)
//		// 记录结果
//		res = append(res, node.Val)
//	}
//	traversal(root)
//	return res
//}

// 迭代法

// 先序遍历：根左右
// 入栈顺序：根右左
// 后序遍历：左右根
// 只需要调整一下入栈顺序为根左右，出栈顺序即为根右左，再将结果反转即可
func postorderTraversal(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	reverse(res)
	return res
}
func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
