package main

// 最小堆 minHeap 的实现
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	// old := *h
	// n := len(old)
	// x := old[n-1]
	// *h = old[:n-1]
	// return x
	x := *h[len(*h)-1]
	*h = *h[:len(*h)-1]
	return x
}

// 最大堆 maxHeap 的实现
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
}

func Constructor() MedianFinder {

}

func (this *MedianFinder) AddNum(num int) {

}

func (this *MedianFinder) FindMedian() float64 {

}
func main() {

}
