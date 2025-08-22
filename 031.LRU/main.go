package main

import "container/list"

type entry struct {
	key   int
	value int
}
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		this.list.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.cache[key]; ok {
		ele.Value.(*entry).value = value
		this.list.MoveToFront(ele)
		return
	}
	if this.list.Len() == this.capacity {
		last := this.list.Back()
		delete(this.cache, last.Value.(*entry).key)
		this.list.Remove(last)
	}
	ele := this.list.PushFront(&entry{key, value})
	this.cache[key] = ele
}
