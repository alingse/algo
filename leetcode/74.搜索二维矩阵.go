/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := m*n - 1
	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/n][mid%n]
		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end

