/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	start := 0
	end := 0
	sum := 0
	minLength := len(nums) + 1
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			minLength = min(minLength, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if minLength == len(nums)+1 {
		return 0
	}
	return minLength
}

// @lc code=end

