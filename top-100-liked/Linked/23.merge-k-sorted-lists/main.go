package main

import "container/heap"

// ListNode 是单链表的定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 创建我们自己的类型，并实现 heap.Interface 接口
// ListNodeHeap 是一个由 ListNode 指针组成的最小堆
type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val } // 这是最小堆的关键
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 和 Pop 需要用指针接收者，因为它们会修改切片的长度
func (h *ListNodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 2. 编写主函数 mergeKLists
func mergeKLists(lists []*ListNode) *ListNode {
	// 创建一个 ListNodeHeap 实例
	h := &ListNodeHeap{}
	heap.Init(h) // 初始化堆

	// 将所有链表的头节点放入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	// 创建结果链表的哨兵节点
	dummyHead := &ListNode{}
	tail := dummyHead

	// 循环直到堆为空
	for h.Len() > 0 {
		// 从堆中取出最小的节点
		minNode := heap.Pop(h).(*ListNode)
		// 连接到结果链表
		tail.Next = minNode
		tail = tail.Next

		// 如果取出的节点还有下一个节点，则将其入堆
		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}

	return dummyHead.Next
}
