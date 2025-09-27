/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1
	for left < right {
		for left < right && !isAlphanumeric(runes[left]) {
			left++
		}
		for left < right && !isAlphanumeric(runes[right]) {
			right--
		}
		if left < right {
			if toLower(runes[left]) != toLower(runes[right]) {
				return false
			}
			left++
			right--
		}
	}
	return true
}

// @lc code=end

