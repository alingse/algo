/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	ans := 0
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			leftMax = max(leftMax, height[left])
			ans += leftMax - height[left]
			left++
		} else {
			rightMax = max(rightMax, height[right])
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

// @lc code=end

