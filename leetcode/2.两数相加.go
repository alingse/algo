/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	v := new(ListNode)
	root := v
	// 进位标志
	pi := 0
	for l1 != nil || l2 != nil || pi > 0 {
		if l1 != nil {
			v.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v.Val += l2.Val
			l2 = l2.Next
		}
		v.Val += pi
		if v.Val >= 10 {
			v.Val -= 10
			pi = 1
		} else {
			pi = 0
		}
		if l1 != nil || l2 != nil || pi > 0 {
			v.Next = new(ListNode)
			v = v.Next
		}
	}
	return root
}

// @lc code=end

