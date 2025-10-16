package main

// 注意边界问题！！！
func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, n-1, 0, m-1
	var res []int
	for top <= bottom && left <= right {
		// 1.左->右
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		// 上边界下移
		top++
		// 2.上->下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		// 右边界左移
		right--
		// 3.右->左
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
		}
		// 下边界上移
		bottom--
		// 4.下->上
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
		}
		// 左边界右移
		left++
	}
	return res
}
