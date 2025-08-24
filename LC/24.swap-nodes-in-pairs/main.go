package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归交换链表节点
// 输入：head: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> null
// 输出：2 -> 1 -> 4 -> 3 -> 6 -> 5 -> 8 -> 7 -> 9 -> null
//func swapPairs(head *ListNode) *ListNode {
//	// 基本情况：链表为空或只有一个节点，无需交换
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	// 获取第二个节点，它将成为交换后的头节点
//	newHead := head.Next
//
//	// head 的 Next 指向递归处理后的子链表
//	head.Next = swapPairs(newHead.Next)
//
//	// 将 newHead 的 Next 指向 head，完成当前两节点的交换
//	newHead.Next = head
//
//	// 返回新的头节点
//	return newHead
//}

// 迭代交换链表节点
// 输入：head: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> null
// 输出：2 -> 1 -> 4 -> 3 -> 6 -> 5 -> 8 -> 7 -> 9 -> null
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // 哨兵节点
	cur := dummy                   // 当前节点
	// 当前节点和下一个节点都存在
	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := cur.Next.Next

		// 交换节点
		first.Next = second.Next
		second.Next = first
		cur.Next = second

		// 移动到下一组
		cur = first
	}

	return dummy.Next
}
