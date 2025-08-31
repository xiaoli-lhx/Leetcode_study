package main

import (
	"container/heap"
	"sort"
)

type Pair struct {
	Number int
	Count  int
}

type IHeap []Pair

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i].Count < h[j].Count }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// 1. 统计频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 创建一个最小堆
	minHeap := &IHeap{}
	heap.Init(minHeap) // 初始化堆

	// 3. 遍历频率，维护一个大小为 k 的最小堆
	for num, count := range freqMap {
		heap.Push(minHeap, Pair{Number: num, Count: count})
		if minHeap.Len() > k {
			heap.Pop(minHeap) // 如果堆大小超过 k，就把最小的那个（堆顶）扔掉
		}
	}

	// 4. 堆里剩下的就是前 k 个高频元素，收集结果
	var result []int
	for minHeap.Len() > 0 {
		// Pop 的结果是 Pair 类型，我们需要它的 Number 字段
		result = append(result, heap.Pop(minHeap).(Pair).Number)
	}
	return result
}

// 另一种实现方式，使用 sort.Slice 进行自定义降序排序
func topKFrequent_Sort(nums []int, k int) []int {
	// 1. 统计频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 将 map 转换为 struct 切片
	type Pair struct {
		Number int
		Count  int
	}
	var pairs []Pair
	for num, count := range freqMap {
		pairs = append(pairs, Pair{Number: num, Count: count})
	}

	// 3. 使用 sort.Slice 进行自定义降序排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	// 4. 取出前 k 个元素
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].Number)
	}
	return result
}
