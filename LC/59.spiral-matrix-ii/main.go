package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1 // 当前数字
	// 四个边界
	top, bottom, left, right := 0, n-1, 0, n-1
	total := n * n // 总数
	for num <= total {
		// 1. 从左到右填充
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++ // 上边界下移
		// 2. 从上到下填充右边界
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right-- // 右边界左移
		// 3. 从右到左填充下边界
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom-- // 下边界上移
		// 4. 从下到上填充左边界
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++ // 左边界右移
	}
	return matrix
}
