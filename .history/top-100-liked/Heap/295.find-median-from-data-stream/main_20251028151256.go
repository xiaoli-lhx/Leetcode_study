package main

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
