/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	// 去除字符串两端的空格
	s = strings.TrimSpace(s)

	// 找到最后一个空格的位置
	lastSpaceIndex := strings.LastIndex(s, " ")

	// 如果没有找到空格，说明整个字符串就是一个单词
	if lastSpaceIndex == -1 {
		return len(s)
	}

	// 返回最后一个单词的长度
	return len(s) - lastSpaceIndex - 1
}

// @lc code=end

