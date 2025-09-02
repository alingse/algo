/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	rs := []rune(s)
	n := len(rs)
	if numRows == 1 || n <= numRows {
		return s
	}
	cycleLen := 2*numRows - 2
	result := make([]rune, 0, n)
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			result = append(result, rs[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				result = append(result, rs[j+cycleLen-i])
			}
		}
	}
	return string(result)
}

// @lc code=end

