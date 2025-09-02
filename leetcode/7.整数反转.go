/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	var result int
	for x != 0 {
		pop := x % 10
		x /= 10
		if result > 214748364 || (result == 214748364 && pop > 7) {
			return 0
		}
		if result < -214748364 || (result == -214748364 && pop < -8) {
			return 0
		}
		result = result*10 + pop
	}
	return result
}

// @lc code=end

