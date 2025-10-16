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

// 递归
func isSymmetric(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}

// 迭代
func isSymmetric1(root *TreeNode) bool {
	queue := []*TreeNode{root.Left, root.Right}
	if root == nil {
		return true
	}
	for len(queue) > 0 {
		leftNode := queue[0]
		rightNode := queue[1]
		queue = queue[2:]
		if leftNode == nil && rightNode == nil {
			// 不能直接返回true 要继续判断
			continue
		}
		if leftNode == nil || rightNode == nil {
			return false
		}
		if leftNode.Val != rightNode.Val {
			return false
		}
		queue = append(queue, leftNode.Left)
		queue = append(queue, rightNode.Right)
		queue = append(queue, leftNode.Right)
		queue = append(queue, rightNode.Left)
	}
	return true
}
