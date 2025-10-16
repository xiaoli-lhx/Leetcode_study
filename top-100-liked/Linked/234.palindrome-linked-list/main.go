package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 后半段链表逆转，然后比较
	// 如何找到后半段链表的头节点？
	// 快慢指针的方式
	fast, slow := head, head
	// 为什么循环终止条件是 fast!= nil && fast.Next!= nil？而不是 fast!= nil || fast.Next!= nil？
	// 只要 fast 或者 fast.Next 中任何一个变成了 nil，循环就应该立即停止，以避免出错。
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// slow 所在处，便是后半段链表的头节点
	// 逆转后半段
	var prev *ListNode
	curr := slow
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	for prev != nil {
		if prev.Val != head.Val {
			return false
		} else {
			prev = prev.Next
			head = head.Next
		}
	}
	return true
}
