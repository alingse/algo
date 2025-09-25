/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rowZero := false
	colZero := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			rowZero = true
			break
		}
	}

	// 使用第一行第一列记录为0的位置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 全部设置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end

