package main

import "container/heap"

// 最小堆 MinHeap 的实现
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 最大堆 MaxHeap 的实现
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 包含两个堆
type MedianFinder struct {
	small *MaxHeap // 存储较小的一半
	large *MinHeap // 存储较大的一半
}

func Constructor() MedianFinder {
	// 初始化两个堆
	small := &MaxHeap{}
	large := &MinHeap{}
	heap.Init(small)
	heap.Init(large)
	return MedianFinder{
		small: small,
		large: large,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 1. 总是将新元素推入 small 最大堆
	heap.Push(this.small, num)
	// 2. 为了维持small中所有元素都<=large中所有元素的特性
	// 	  我们将small中最大的元素(即堆顶)弹出，并推入 large 最小堆
}

func (this *MedianFinder) FindMedian() float64 {

}
func main() {

}
