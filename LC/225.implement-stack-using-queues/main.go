package main

import "container/list"

// 方法一，pop复杂

type MyStack struct {
	queue *list.List
}

func Constructor() MyStack {
	return MyStack{
		queue: list.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.queue.PushBack(x)
}

func (this *MyStack) Pop() int {
	n := this.queue.Len()
	for i := 0; i < n-1; i++ {
		n := this.queue.Front().Value.(int)
		this.queue.Remove(this.queue.Front())
		this.Push(n)
	}
	x := this.queue.Front().Value.(int)
	this.queue.Remove(this.queue.Front())
	return x
}

func (this *MyStack) Top() int {
	x := this.Pop()
	this.Push(x)
	return x
}

func (this *MyStack) Empty() bool {
	return this.queue.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// 方法二，push复杂

//type MyStack struct {
//	queue *list.List
//}
//
//
//func Constructor() MyStack {
//	return MyStack{
//		queue: list.New(),
//	}
//}
//
//
//func (this *MyStack) Push(x int)  {
//	i:=this.queue.PushBack(x)
//	this.queue.MoveToFront(i)
//}
//
//
//func (this *MyStack) Pop() int {
//	front:=this.queue.Front()
//	this.queue.Remove(front)
//	return front.Value.(int)
//}
//
//
//func (this *MyStack) Top() int {
//	return this.queue.Front().Value.(int)
//}
//
//
//func (this *MyStack) Empty() bool {
//	return this.queue.Len()==0
//}
//
//
/**
* Your MyStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.Empty();
 */
