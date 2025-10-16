package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归法 (自顶向下)：先把一个长链表不断切分，直到只剩单个节点，然后再一层层合并回来。
// 时间复杂度O(n log n) 空间复杂度O(log n)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针找寻链表中点
	fast, slow := head, head
	// 中点的前一个指针
	var slowPrev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slowPrev = slow
		slow = slow.Next
	}
	// 拆分
	slowPrev.Next = nil
	// 排序
	newHead := slow
	head = sortList(head)
	newHead = sortList(newHead)
	return merge(head, newHead)
}
func merge(head, newHead *ListNode) *ListNode {
	// 合并
	dummy := &ListNode{}
	curr := dummy
	for head != nil && newHead != nil {
		if head.Val < newHead.Val {
			curr.Next = head
			head = head.Next
		} else {
			curr.Next = newHead
			newHead = newHead.Next
		}
		curr = curr.Next
	}
	if head != nil {
		curr.Next = head
	} else {
		curr.Next = newHead
	}
	return dummy.Next
}

// TODO: 常数级空间复杂度的排序算法
// 迭代法 (自底向上)：它把整个链表一开始就看作是 N 个已经排好序的、长度为 1 的子链表。然后，我们开始合并它们。
// 时间复杂度O(n log n) 空间复杂度O(1)
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 计算链表总长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummy := &ListNode{Next: head}
	// size 从1开始，每次翻倍
	for size := 1; size < length; size <<= 1 {
		// prev 是已排序部分的尾节点，curr 是待处理部分的头节点
		prev, curr := dummy, dummy.Next
		// 内层循环
		for curr != nil {
			head1 := curr
			for i := 1; i < size && curr.Next != nil; i++ {
				curr = curr.Next
			}
			head2 := curr.Next
			curr.Next = nil
			curr = head2
			for i := 1; i < size && curr != nil && curr.Next != nil; i++ {
				curr = curr.Next
			}
			var next *ListNode
			if curr != nil {
				next = curr.Next
				curr.Next = nil
			}
			merged := merge(head1, head2)
			prev.Next = merged
			for prev.Next != nil {
				prev = prev.Next
			}
			curr = next
		}
	}
	return dummy.Next
}
