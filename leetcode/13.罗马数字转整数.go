/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	romanNumerals := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := 0
	for {
		for _, numeral := range romanNumerals {
			if strings.HasPrefix(s, numeral.symbol) {
				result += numeral.value
				s = s[len(numeral.symbol):] // 移除已处理的符号
			}
		}
		if s == "" {
			break // 如果字符串已空，结束循环
		}
	}

	return result
}

// @lc code=end

