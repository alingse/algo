/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	chars := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if check(board, chars, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func check(board [][]byte, chars []byte, i, j, k int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != chars[k] {
		return false
	}
	if k == len(chars)-1 {
		return true
	}
	origin := board[i][j]
	board[i][j] = byte(0)
	res := check(board, chars, i+1, j, k+1) || check(board, chars, i-1, j, k+1) || check(board, chars, i, j+1, k+1) || check(board, chars, i, j-1, k+1)
	board[i][j] = origin
	return res
}

// @lc code=end

