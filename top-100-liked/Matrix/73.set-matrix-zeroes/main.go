package main

// 时间复杂度O(mn)
// 空间复杂度O(m+n)
//func setZeroes(matrix [][]int) {
//	rows := make([]bool, len(matrix))
//	cols := make([]bool, len(matrix[0]))
//	for i, r := range matrix {
//		for j, v := range r {
//			if v == 0 {
//				rows[i] = true
//				cols[j] = true
//			}
//		}
//	}
//	for i, r := range matrix {
//		for j := range r {
//			if rows[i] || cols[j] {
//				r[j] = 0
//			}
//		}
//	}
//}

// 标记变量法
// 时间复杂度O(mn)
// 空间复杂度O(1)
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row, col := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col = true
			break
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
}
