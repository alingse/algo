/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	n := len(s)
	i := 0
	// 1. 去除前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}
	// 2. 处理符号
	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}
	// 3. 转换数字部分
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		// 检查溢出
		if result > (2147483647-digit)/10 {
			if sign == 1 {
				return 2147483647
			} else {
				return -2147483648
			}
		}
		result = result*10 + digit
		i++
	}
	return result * sign
}

// @lc code=end

