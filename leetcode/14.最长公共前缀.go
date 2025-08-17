/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:] {
		if len(s) < len(prefix) {
			prefix = prefix[:len(s)]
		}
		for j := 0; j < len(prefix) && j < len(s); j++ {
			if prefix[j] != s[j] {
				prefix = prefix[:j]
				break
			}
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

// @lc code=end

