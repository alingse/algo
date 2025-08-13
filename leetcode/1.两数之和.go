/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok {
			return []int{j, i}
		}
		numsMap[num] = i
	}
	// 如果没有找到，返回一个空的切片
	return []int{}
}

// @lc code=end

