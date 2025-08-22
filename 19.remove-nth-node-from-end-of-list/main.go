package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// dummy->1->2->3->4->5->null, n = 2
// for cur.Next != nil 统计链表长度 count = 5,n = 2
// count-n = 3
// for i := 0; i < count-n; i++ 找到倒数第n个节点
// cur=dummy i=0 cur.Next=1
// cur=1 i=1 cur.Next=2 cur=2
// cur=2 i=2 cur.Next=3 cur=3
// cur=3 i=3 cur.Next=4
// cur=4 i=4 cur.Next=5
// cur.Next = cur.Next.Next 删除倒数第n个节点
// 1->2->3->5->null
//
//	func removeNthFromEnd(head *ListNode, n int) *ListNode {
//		dummy := &ListNode{Next: head} // 虚拟头节点
//		cur := dummy
//		count := 0
//		for cur.Next != nil {
//			cur = cur.Next
//			count++
//		}
//		cur = dummy
//		for i := 0; i < count-n; i++ {
//			cur = cur.Next
//		}
//		cur.Next = cur.Next.Next
//		return dummy.Next
//	}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // 虚拟头节点
	fast, slow := dummy, dummy     // 快慢指针
	// 快指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// slow和fast同时走，直至fast到达链表尾部
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 此时slow位于待删除节点的前一个节点
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 通过在 fast 和 slow 之间建立一个固定的距离，将“找到倒数第 n 个节点”这个任务，转化为“在 fast 走完整个链表时，slow 位于正确位置
// dummy->1->2->3->4->5->null, n=2
// fast = dummy fast.Next = 1 fast = 1 i=0
// fast = 1 fast.Next = 2 fast = 2 i=1
