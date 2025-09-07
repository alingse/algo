/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	i := len(digits) - 1
	for i >= 0 {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
		i--
	}
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}

// @lc code=end

