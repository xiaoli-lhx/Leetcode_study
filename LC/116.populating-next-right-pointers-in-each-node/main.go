package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 空间复杂度 O(W),其中 W 是树的最大宽度（在完美二叉树中大约是 N/2 个节点）。
//func connect(root *Node) *Node {
//	if root == nil {
//		return nil
//	}
//	queue := []*Node{root}
//	for len(queue) > 0 {
//		size := len(queue)
//		for i := 0; i < size; i++ {
//			node := queue[0]
//			queue = queue[1:]
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//			if i < size-1 {
//				node.Next = queue[0]
//			}
//		}
//	}
//	return root
//}

// 只使用常量级额外空间
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	levelStart := root
	for levelStart.Left != nil {
		current := levelStart
		for current != nil {
			current.Left.Next = current.Right
			if current.Next != nil {
				current.Right.Next = current.Next.Left
			}
			current = current.Next
		}
		levelStart = levelStart.Left
	}
	return root
}
