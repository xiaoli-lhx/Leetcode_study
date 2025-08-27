package main

// 双向链表模拟栈
// import "container/list"
//
//	func removeDuplicates(s string) string {
//		stack := list.New()
//		stack.PushBack(s[0])
//		for i := 1; i < len(s); i++ {
//			if s[i] == stack.Back().Value.(byte) {
//				stack.Remove(stack.Back())
//			} else {
//				stack.PushBack(s[i])
//			}
//		}
//		result := ""
//		for e := stack.Front(); e != nil; e = e.Next() {
//			result += e.Value.(string)
//		}
//		return result
//	}
//
// 切片模拟栈
func removeDuplicates(s string) string {
	var stack = []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			// pop
			stack = stack[:len(stack)-1]
		} else {
			// push
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
