package main

func isHappy(n int) bool {
	seen := make(map[int]bool)
	// 循环结束条件
	// 1. n==1 说明已经是1了
	// 2. seen[n] 说明已经计算过了
	for n != 1 && !seen[n] {
		seen[n] = true
		// 计算下一个值
		n = getNext(n)
	}
	return n == 1
}
func getNext(n int) int {
	sum := 0
	for n > 0 {
		// 取最后一位
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
