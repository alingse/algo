/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	start := 0
	match := 0
	for i := 0; i <= n-m; i++ {
		start = i
		match = 0
		for match < m {
			if haystack[start] == needle[match] {
				start++
				match++
				continue
			}
			break
		}
		if match == m {
			return i
		}
	}
	return -1
}

// @lc code=end

