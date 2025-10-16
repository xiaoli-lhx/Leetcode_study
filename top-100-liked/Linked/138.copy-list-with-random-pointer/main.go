package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 哈希表
// 空间复杂度 O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	hashMap := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		// 创建新的链表
		hashMap[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		newNode := hashMap[curr]
		newNode.Next = hashMap[curr.Next]
		newNode.Random = hashMap[curr.Random]
		curr = curr.Next
	}
	return hashMap[head]
}
