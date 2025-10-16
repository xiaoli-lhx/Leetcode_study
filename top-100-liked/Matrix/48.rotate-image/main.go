package main

import "fmt"

// 沿对角线翻转再倒序

//func rotate(matrix [][]int) {
//	n := len(matrix)
//	// 沿对角线翻转
//	for i := 0; i < n; i++ {
//		// 只处理对角线下的元素
//		for j := 0; j < i; j++ {
//			// fmt.Print(matrix[i][j])
//			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
//		}
//	}
//	// 倒序
//	for i := 0; i < n; i++ {
//		left, right := 0, n-1
//		for left <= right {
//			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
//			left++
//			right--
//		}
//	}
//}

// 原地旋转
func rotate(matrix [][]int) {
	// 分块
	// (row,col) 顺时针旋转90° -> newRow=col,newCol=n-1-row
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			// 左上块   左下块              右下块              右上块
			matrix[i][j], matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i] =
				matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}
