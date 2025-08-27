package main

import "container/list"

func isValid(s string) bool {
	paris := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	list := list.New()
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			// push
			list.PushBack(ch)
		}
		if ch == ')' || ch == '}' || ch == ']' {
			if list.Len() == 0 {
				return false
			} else {
				// pop
				old := list.Back()
				list.Remove(old)
				if old.Value != paris[ch] {
					return false
				}
			}
		}
	}
	if list.Len() != 0 {
		return false
	}
	return true
}
