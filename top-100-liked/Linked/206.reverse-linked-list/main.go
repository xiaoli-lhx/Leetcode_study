package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代法

//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode // prev指向反转后的新链表，初始值为nil
//	curr := head       // curr指向待反转的链表
//	for curr != nil {
//		nextTmp := curr.Next // nextTmp指向curr的下一个节点
//		curr.Next = prev     // curr的下一个节点指向prev
//		prev = curr          // prev指向curr
//		curr = nextTmp       // curr指向nextTmp
//	}
//	return prev
//}

//1->2->3->4->5->nil
//curr = 1, nextTmp = 2, prev = nil
//curr.Next = nil, prev = 1, curr = 2
//curr = 2, nextTmp = 3, prev = 1
//curr.Next = 1, prev = 2, curr = 3
//curr = 3, nextTmp = 4, prev = 2
//curr.Next = 2, prev = 3, curr = 4
//curr = 4, nextTmp = 5, prev = 3
//curr.Next = 3, prev = 4, curr = 5
//curr = 5, nextTmp = nil, prev = 4
//prev = 5, curr = nil
//5->4->3->2->1->nil

// 递归法

func reverseList(head *ListNode) *ListNode {
	// 递归终止条件
	// 1. 如果 head 为 nil，说明链表为空，直接返回 nil
	// 2. 如果 head.Next 为 nil，说明链表只有一个节点，本身就是反转的，直接返回 head
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
