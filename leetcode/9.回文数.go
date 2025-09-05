/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x != 0 && x%10 == 0 {
		return false
	}
	revertedNumber := 0
	// 依次从 x 的末尾开始取数据，每个乘以10，进入下一轮，最后得到
	// 123321 --> (((1+0)*10 + 2)*10 + 3)*10
	// 直到 x 是 123 ，而同时 revertedNumber 也是 123
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	// 如果是 1234321 的时候
	// x 是 123 而 revertedNumber 是 1234, 这个时候末尾的那个位数就可以忽略(中间位置)
	return x == revertedNumber || x == revertedNumber/10
}

// @lc code=end

