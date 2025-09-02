/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	needles := []rune(needle)
	next := make([]int, m)
	// 1. i 从 1 开始遍历, j 从 -1 开始
	// 2. next[0] = -1
	// 3. 使用 next 回溯 j 直到 j = -1 或 needles[i] == needles[j+1]
	// 4. 如果 needles[i] == needles[j+1], 则 j++
	// 5. 将 j 赋值给 next[i]
	// 6. i++
	i := 1
	j := -1
	next[0] = -1
	for i < m-1 {
		for needles[i] != needles[j+1] && j >= 0 {
			j = next[j]
		}
		if needles[i] == needles[j+1] {
			j++
		}
		next[i] = j
		i++
	}
	haystacks := []rune(haystack)
	i = 0
	j = -1
	// 1. i 从 0 开始遍历, j 从 -1 开始
	// 2. 使用 next 回溯 j 直到 j = -1 或 haystacks[i] == needles[j+1]
	// 3. 如果 haystacks[i] == needles[j+1], 则 j++
	// 4. 如果 j == m - 1, 则找到了匹配项, 返回 i - m + 1 (说明累积匹配了 m 个字符)
	// 5. i++
	for i < n {
		for haystacks[i] != needles[j+1] && j >= 0 {
			j = next[j]
		}
		if haystacks[i] == needles[j+1] {
			j++
		}
		if j == m-1 {
			return i - m + 1
		}
		i++
	}
	return -1
}

// @lc code=end

