/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	totalLeft := (m + n + 1) / 2
	
	for left <= right {
		// nums1 的分割点
		i := (left + right) / 2
		// nums2 的分割点
		j := totalLeft - i
		
		// 处理边界情况
		nums1Left := -1 << 31
		if i > 0 {
			nums1Left = nums1[i-1]
		}
		
		nums1Right := 1<<31 - 1
		if i < m {
			nums1Right = nums1[i]
		}
		
		nums2Left := -1 << 31
		if j > 0 {
			nums2Left = nums2[j-1]
		}
		
		nums2Right := 1<<31 - 1
		if j < n {
			nums2Right = nums2[j]
		}
		
		// 检查是否找到正确的分割点
		if nums1Left <= nums2Right && nums2Left <= nums1Right {
			// 找到了中位数
			if (m+n)%2 == 1 {
				return float64(max(nums1Left, nums2Left))
			} else {
				return float64(max(nums1Left, nums2Left)+min(nums1Right, nums2Right)) / 2.0
			}
		} else if nums1Left > nums2Right {
			// i 需要向左移动
			right = i - 1
		} else {
			// i 需要向右移动
			left = i + 1
		}
	}
	
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

