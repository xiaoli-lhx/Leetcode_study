package main

// 单链表实现

type ListNode struct {
	Val  int
	Next *ListNode
}
type MyLinkedList struct {
	dummyHead *ListNode // 虚拟头结点
	size      int       // 链表长度
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &ListNode{},
		size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	current := this.dummyHead.Next
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	newNode := &ListNode{Val: val}
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	newNode.Next = prev.Next
	prev.Next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	this.size--
}
