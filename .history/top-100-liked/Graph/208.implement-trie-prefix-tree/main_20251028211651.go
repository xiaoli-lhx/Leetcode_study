package main

type Trie struct {
	// children 存储指向26个小写字母的子节点指针
	// children[0] -> 'a'  children[1] -> 'b'
	children [26]*Trie
	// isEnd 标记从根到当前节点是否形成一个完整的单词
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}
