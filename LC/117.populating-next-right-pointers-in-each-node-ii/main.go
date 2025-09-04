package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	levelStart := root
	for levelStart != nil {
		nextLevelStart := &Node{}
		tail := nextLevelStart
		current := levelStart
		for current != nil {
			if current.Left != nil {
				tail.Next = current.Left
				tail = tail.Next
			}
			if current.Right != nil {
				tail.Next = current.Right
				tail = tail.Next
			}
			current = current.Next
		}
		levelStart = nextLevelStart.Next
	}
	return root
}
